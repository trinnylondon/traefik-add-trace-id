package platform

import (
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
		name string
	}{
		{name: "no trace id"},
		{name: "empty trace id"},
		{name: "custom name"},
		{name: "no traceid no prefix"},
		{name: "custom traceid and prefix"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
