package main

import (
	"context"
	"fmt"
	"net/http"
)

// Config tmp
type Config struct{}

// CreateConfig tmp
func CreateConfig() *Config { return &Config{} }

// HT tmp
type HT struct {
	next http.Handler
	name string
}

// New tmp
func New(_ context.Context, next http.Handler, _ *Config, name string) (http.Handler, error) {
	fmt.Println("###########################################")
	fmt.Println("HT plugin loaded")
	return &HT{
		next: next,
		name: name,
	}, nil
}

func (e *HT) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Del("Date")
	req.Header.Del("Date")

	e.next.ServeHTTP(rw, req)
}
