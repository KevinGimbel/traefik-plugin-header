// Package plugin_header a plugin to adds default headers
package plugin_header

import (
	"context"
	"net/http"
)

// Config holds the plugin configuration.
type Config struct {
	Headers []Header `json:"headers,omitempty"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// CreateConfig creates and initializes the plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

type defaultHeaders struct {
	name    string
	next    http.Handler
	headers []Header
}

// New creates and returns a plugin instance.
func New(_ context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &defaultHeaders{
		name:    name,
		next:    next,
		headers: config.Headers,
	}, nil
}

func (h *defaultHeaders) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	for _, h := range h.headers {
		rw.Header().Add(h.Key, h.Value)
	}

	h.next.ServeHTTP(rw, req)
}
