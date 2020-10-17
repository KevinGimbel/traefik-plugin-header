// Package plugin_header a plugin to adds default headers
package traefik_plugin_header

import (
	"context"
	"net/http"
)

// Config holds the plugin configuration.
type Config struct {
	Headers []Header `json:"headers,omitempty"`
}

// Header represents a HTTP header config
type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// CreateConfig creates and initializes the plugin configuration.
func CreateConfig() *Config {
	return &Config{Headers: []Header{}}
}

type headerHTTPHandler struct {
	name    string
	next    http.Handler
	headers []Header
}

// New creates and returns a plugin instance.
func New(_ context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &headerHTTPHandler{
		name:    name,
		next:    next,
		headers: config.Headers,
	}, nil
}

func (h *headerHTTPHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	for _, h := range h.headers {
		rw.Header().Add(h.Key, h.Value)
	}

	h.next.ServeHTTP(rw, req)
}
