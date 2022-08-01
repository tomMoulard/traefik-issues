package tcpinflightreq

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v2/pkg/config/dynamic"
	"github.com/traefik/traefik/v2/pkg/tcp"
)

func TestNewInFlightReq(t *testing.T) {
	testCases := []struct {
		desc   string
		config dynamic.TCPInFlightReq
	}{
		{
			desc:   "Empty config",
			config: dynamic.TCPInFlightReq{},
		},
		{
			desc: "Valid config",
			config: dynamic.TCPInFlightReq{
				Amount: 0,
			},
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			next := tcp.HandlerFunc(func(conn tcp.WriteCloser) {})
			inFlightReq, err := New(context.Background(), next, test.config, "traefik_inFlightReq_Test")

			require.NoError(t, err)
			assert.NotNil(t, inFlightReq)
		})
	}
}

func TestInFlightReq_ServeTCP(t *testing.T) {
	testCases := []struct {
		desc               string
		config             dynamic.TCPInFlightReq
		concurrentRequests int
		expectError        bool
	}{
		{
			desc:               "One request, one allowed",
			concurrentRequests: 1,
			config:             dynamic.TCPInFlightReq{Amount: 1},
		},
		{
			desc:               "More request than ones allowed",
			concurrentRequests: 42,
			expectError:        true,
		},
		{
			desc:               "Less request than ones allowed",
			concurrentRequests: 42,
			config:             dynamic.TCPInFlightReq{Amount: 4242},
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			var wg sync.WaitGroup
			wg.Add(test.concurrentRequests)
			var nbReq int64
			next := tcp.HandlerFunc(func(conn tcp.WriteCloser) {
				conn.Close()
				wg.Done()
				atomic.AddInt64(&nbReq, 1)
			})

			inFlightReq, err := New(context.Background(), next, test.config, "traefikTest")
			require.NoError(t, err)

			server, client := net.Pipe()

			for i := 0; i < test.concurrentRequests; i++ {
				go func() {
					inFlightReq.ServeTCP(&contextWriteCloser{client, addr{fmt.Sprintf("127.0.0.1:%d", i)}})
				}()
			}

			_, err = ioutil.ReadAll(server)
			require.NoError(t, err)

			wg.Wait()
			t.Log(test, nbReq)
			if test.expectError {
				assert.Less(t, int(nbReq), test.concurrentRequests)
			} else {
				assert.Equal(t, int(nbReq), test.concurrentRequests)
			}
		})
	}
}

type contextWriteCloser struct {
	net.Conn
	addr
}

type addr struct {
	remoteAddr string
}

func (a addr) Network() string {
	panic("implement me")
}

func (a addr) String() string {
	return a.remoteAddr
}

func (c contextWriteCloser) CloseWrite() error {
	panic("implement me")
}

func (c contextWriteCloser) RemoteAddr() net.Addr { return c.addr }

func (c contextWriteCloser) Context() context.Context {
	return context.Background()
}
