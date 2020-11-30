package platform

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServeHTTP(t *testing.T) {
	// cfg := plugindemo.CreateConfig()

	// ctx := context.Background()
	// next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	// handler, err := plugindemo.New(ctx, next, cfg, "demo-plugin")
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// recorder := httptest.NewRecorder()

	// req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// handler.ServeHTTP(recorder, req)

	tests := []struct {
		name       string
		config     *Config
		assertFunc func(t *testing.T) http.Handler
	}{
		{
			name:   "no trace id",
			config: &Config{},
			assertFunc: func(t *testing.T) http.Handler {
				t.Helper()
				return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					if getTraceIdHeader(t, req, "X-Trace-Id") == "" {
						t.Fatalf("got empty traceId header for %+v", req.Header)
					}
				})
			},
		},
		{
			name: "custom name",
			config: &Config{
				HeaderName: "Other-Name",
			},
			assertFunc: func(t *testing.T) http.Handler {
				t.Helper()
				return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					if getTraceIdHeader(t, req, "Other-Name") == "" {
						t.Fatalf("got empty traceId header for %+v", req.Header)
					}
				})
			},
		},
		{
			name: "with prefix",
			config: &Config{
				HeaderPrefix: "myorg",
			},
			assertFunc: func(t *testing.T) http.Handler {
				t.Helper()
				return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					hdr := getTraceIdHeader(t, req, "X-Trace-Id")

					if !strings.HasPrefix(hdr, "myorg:") {
						t.Fatalf("no prefix in traceId: %+v", req.Header["X-Trace-Id"])
					}
				})
			},
		},
		{
			name: "custom traceid and prefix",
			config: &Config{
				HeaderPrefix: "myorg",
				HeaderName:   "Other-Name",
			},
			assertFunc: func(t *testing.T) http.Handler {
				t.Helper()
				return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
					hdr := getTraceIdHeader(t, req, "Other-Name")

					if !strings.HasPrefix(hdr, "myorg:") {
						t.Fatalf("no prefix in traceId: %+v", req.Header["Other-Name"])
					}
				})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			handler, err := New(ctx, tt.assertFunc(t), tt.config, "trace-id-test")
			if err != nil {
				t.Fatalf("error with new redirect: %+v", err)
			}
			recorder := httptest.NewRecorder()
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost/", nil)
			if err != nil {
				t.Fatalf("error with new request: %+v", err)
			}

			handler.ServeHTTP(recorder, req)

		})
	}
}

func getTraceIdHeader(t *testing.T, req *http.Request, headerName string) string {
	t.Helper()
	headerArr := req.Header[headerName]
	if len(headerArr) == 1 {
		return headerArr[0]
	}
	return ""
}
