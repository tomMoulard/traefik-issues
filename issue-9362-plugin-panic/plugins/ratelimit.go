package ratelimit

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/mailgun/holster/v4/clock"
	"github.com/mailgun/holster/v4/collections"
	"golang.org/x/time/rate"
)

const (
	maxSources = 65536
)

type Checker struct {
	authorizedIPs    []*net.IP
	authorizedIPsNet []*net.IPNet
}

type Config struct {
	ExcludedIps       []string `json:"excludedIps,omitempty"`
	IpFromHeaderName  string   `json:"ipFromHeaderName,omitempty"`
	UidFromHeaderName string   `json:"uidFromHeaderName,omitempty"`
	Average           int64    `json:"average,omitempty"`
	Period            int      `json:"period,omitempty"`
	Burst             int64    `json:"burst,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

type RateLimiter struct {
	name              string
	rate              rate.Limit // reqs/s
	burst             int64
	maxDelay          time.Duration
	ttl               int
	next              http.Handler
	buckets           *collections.TTLMap
	checker           *Checker
	ipFromHeaderName  string
	uidFromHeaderName string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	buckets := collections.NewTTLMap(maxSources)
	clock.Freeze(time.Now())

	burst := config.Burst
	if burst < 1 {
		burst = 1
	}

	period := time.Duration(config.Period) * time.Second
	if period < 0 {
		return nil, fmt.Errorf("negative value not valid for period: %v", period)
	}
	if period == 0 {
		period = time.Second
	}

	// if config.Average == 0, in that case,
	// the value of maxDelay does not matter since the reservation will (buggily) give us a delay of 0 anyway.
	var maxDelay time.Duration
	var rtl float64
	if config.Average > 0 {
		rtl = float64(config.Average*int64(time.Second)) / float64(period)
		// maxDelay does not scale well for rates below 1,
		// so we just cap it to the corresponding value, i.e. 0.5s, in order to keep the effective rate predictable.
		// One alternative would be to switch to a no-reservation mode (Allow() method) whenever we are in such a low rate regime.
		if rtl < 1 {
			maxDelay = 500 * time.Millisecond
		} else {
			maxDelay = time.Second / (time.Duration(rtl) * 2)
		}
	}
	// Make the ttl inversely proportional to how often a rate limiter is supposed to see any activity (when maxed out),
	// for low rate limiters.
	// Otherwise, just make it a second for all the high rate limiters.
	// Add an extra second in both cases for continuity between the two cases.
	ttl := 1
	if rtl >= 1 {
		ttl++
	} else if rtl > 0 {
		ttl += int(1 / rtl)
	}

	checker := &Checker{}
	if config.ExcludedIps != nil || len(config.ExcludedIps) > 0 {
		checker, _ = NewChecker(config.ExcludedIps)
	}

	return &RateLimiter{
		name:              name,
		rate:              rate.Limit(rtl),
		burst:             burst,
		maxDelay:          maxDelay,
		ttl:               ttl,
		next:              next,
		buckets:           buckets,
		checker:           checker,
		ipFromHeaderName:  config.IpFromHeaderName,
		uidFromHeaderName: config.UidFromHeaderName,
	}, nil
}

func GetIP(req *http.Request, requestHeaderName string) string {
	ipAddr := req.Header.Get(requestHeaderName)
	ip, _, err := net.SplitHostPort(ipAddr)
	if err != nil {
		return ipAddr
	}
	if ip == "" {
		ip, _, err := net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			return req.RemoteAddr
		}
		return ip
	}
	return ip
}

func (rl *RateLimiter) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	reqIp := GetIP(req, rl.ipFromHeaderName)
	if contain, _ := rl.checker.Contains(reqIp); contain {
		rl.next.ServeHTTP(rw, req)
	} else {
		var bucket *rate.Limiter
		uid := req.Header.Get(rl.uidFromHeaderName)
		source := fmt.Sprintf("%s%s", reqIp, req.URL.EscapedPath())
		if uid != "" {
			source = fmt.Sprintf("%s%s", uid, req.URL.EscapedPath())
		}

		if rlSource, exists := rl.buckets.Get(source); exists {
			bucket = rlSource.(*rate.Limiter)
		} else {
			bucket = rate.NewLimiter(rl.rate, int(rl.burst))
		}

		if err := rl.buckets.Set(source, bucket, rl.ttl); err != nil {
			fmt.Printf("could not insert/update bucket: %v\n", err)
			http.Error(rw, "could not insert/update bucket", http.StatusInternalServerError)
			return
		}

		res := bucket.Reserve()
		if !res.OK() {
			http.Error(rw, "No bursty traffic allowed", http.StatusTooManyRequests)
			return
		}
		delay := res.Delay()

		fmt.Println("RateLimited: #######################################################################################")
		fmt.Printf("RateLimited: source: %s, rl.rate: %v, burst=%d, res.ok=%v, res.delay=%v, "+
			"maxDelay=%v, ttl=%v\n", source, rl.rate, int(rl.burst), res.OK(), delay, rl.maxDelay, rl.ttl)
		if delay > rl.maxDelay {
			res.Cancel()
			rl.serveDelayError(rw, req, delay)
			return
		}
		time.Sleep(delay)
		rl.next.ServeHTTP(rw, req)
	}
}

func (rl *RateLimiter) serveDelayError(w http.ResponseWriter, r *http.Request, delay time.Duration) {
	w.Header().Add("Retry-After", fmt.Sprintf("%.0f", delay.Seconds()))
	w.Header().Add("X-Retry-In", delay.String())
	w.WriteHeader(http.StatusTooManyRequests)

	if _, err := w.Write([]byte(http.StatusText(http.StatusTooManyRequests))); err != nil {
		fmt.Printf("could not serve 429: %v\n", err)
	}
}

func NewChecker(trustedIPs []string) (*Checker, error) {
	if len(trustedIPs) == 0 {
		return nil, errors.New("no trusted IPs provided")
	}

	checker := &Checker{}

	for _, ipMask := range trustedIPs {
		if ipAddr := net.ParseIP(ipMask); ipAddr != nil {
			checker.authorizedIPs = append(checker.authorizedIPs, &ipAddr)
			continue
		}

		_, ipAddr, err := net.ParseCIDR(ipMask)
		if err != nil {
			return nil, fmt.Errorf("parsing CIDR trusted IPs %s: %w", ipAddr, err)
		}
		checker.authorizedIPsNet = append(checker.authorizedIPsNet, ipAddr)
	}

	return checker, nil
}

func (c *Checker) Contains(addr string) (bool, error) {
	if len(addr) == 0 {
		return false, errors.New("empty IP address")
	}

	ipAddr, err := parseIP(addr)
	if err != nil {
		return false, fmt.Errorf("unable to parse address: %s: %w", addr, err)
	}

	return c.ContainsIP(ipAddr), nil
}

func (c *Checker) ContainsIP(addr net.IP) bool {
	for _, authorizedIP := range c.authorizedIPs {
		if authorizedIP.Equal(addr) {
			return true
		}
	}

	for _, authorizedNet := range c.authorizedIPsNet {
		if authorizedNet.Contains(addr) {
			return true
		}
	}

	return false
}

func (c *Checker) IsAuthorized(addr string) error {
	var invalidMatches []string

	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		host = addr
	}

	ok, err := c.Contains(host)
	if err != nil {
		return err
	}

	if !ok {
		invalidMatches = append(invalidMatches, addr)
		return fmt.Errorf("%q matched none of the trusted IPs", strings.Join(invalidMatches, ", "))
	}

	return nil
}

func parseIP(addr string) (net.IP, error) {
	ip := net.ParseIP(addr)
	if ip == nil {
		return nil, fmt.Errorf("can't parse IP from address %s", addr)
	}

	return ip, nil
}
