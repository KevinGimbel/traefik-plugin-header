package traefik_plugin_header

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		desc    string
		headers []Header
		expErr  bool
	}{
		{
			desc:    "should return no error",
			headers: []Header{{Key: "hello", Value: "world"}},
			expErr:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			cfg := &Config{
				Headers: test.headers,
			}

			if _, err := New(context.Background(), nil, cfg, "name"); test.expErr && err == nil {
				t.Errorf("expected error on bad regexp format")
			}
		})
	}
}

func TestServeHTTP(t *testing.T) {
	tests := []struct {
		desc          string
		headers       []Header
		expHeader     []Header
		expNextCall   bool
		expStatusCode int
	}{
		{
			desc:          "should add default header",
			headers:       []Header{{Key: "hello", Value: "world"}},
			expHeader:     []Header{{Key: "hello", Value: "world"}},
			expNextCall:   true,
			expStatusCode: http.StatusOK,
		}, {
			desc:          "should add default header with space in value",
			headers:       []Header{{Key: "Accept-Type", Value: "text/plain text/javascript"}},
			expHeader:     []Header{{Key: "Accept-Type", Value: "text/plain text/javascript"}},
			expNextCall:   true,
			expStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			cfg := &Config{
				Headers: test.headers,
			}

			nextCall := false
			next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				nextCall = true
			})

			handler, err := New(context.Background(), next, cfg, "headers")
			if err != nil {
				t.Fatal(err)
			}

			recorder := httptest.NewRecorder()

			req := httptest.NewRequest(http.MethodGet, "http://localhost", nil)

			handler.ServeHTTP(recorder, req)

			if nextCall != test.expNextCall {
				t.Errorf("next handler should not be called")
			}

			for _, header := range test.expHeader {
				if got := recorder.Result().Header.Get(header.Key); got != header.Value {
					t.Errorf("expected value %s for header %s but got %s", header.Value, header.Key, got)
				}
			}
			if recorder.Result().StatusCode != test.expStatusCode {
				t.Errorf("got status code %d, want %d", recorder.Code, test.expStatusCode)
			}
		})
	}
}
