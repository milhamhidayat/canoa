package test

import (
	"net/http"
	"testing"
)

// RoundTripFunc is a func type to perform executing single HTTP transaction
// Mock http.Client by replacing http.Transport
// Source: http://hassansin.github.io/Unit-Testing-http-client-in-Go#2-by-replacing-httptransport
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip execute a single HTTP transaction
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewClient returns *http.Client with Transport replaced to avoid making real calls
func NewClient(t *testing.T, fn RoundTripFunc) *http.Client {
	t.Helper()

	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

// FuncCall is helper to test function call
type FuncCall struct {
	Called bool
	Input  []interface{}
	Output []interface{}
}
