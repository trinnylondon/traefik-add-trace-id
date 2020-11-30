package platform

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

const defaultTraceId = "X-Trace-Id"

// Config the plugin configuration.
type Config struct {
	HeaderPrefix string `json:"headerPrefix"`
	HeaderName   string `json:"headerName"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		HeaderPrefix: "",
		HeaderName:   defaultTraceId,
	}
}

// TraceIDHeader header if it's missing
type TraceIDHeader struct {
	headerName   string
	headerPrefix string
	name         string
	next         http.Handler
}

// New created a new TraceIDHeader plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	tIDHdr := &TraceIDHeader{
		next: next,
		name: name,
	}

	if config == nil {
		return nil, fmt.Errorf("config can not be nil")
	}

	if config.HeaderName == "" {
		tIDHdr.headerName = defaultTraceId
	} else {
		tIDHdr.headerName = config.HeaderName
	}

	if config.HeaderPrefix != "" {
		tIDHdr.headerPrefix = fmt.Sprintf("%s:", config.HeaderPrefix)
	}

	return tIDHdr, nil

}

func (t *TraceIDHeader) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	headerArr := req.Header[t.headerName]
	randomUUID := fmt.Sprintf("%s%s", t.headerPrefix, uuid.New().String())
	if len(headerArr) == 0 {
		req.Header.Add(t.headerName, randomUUID)
	} else if headerArr[0] == "" {
		req.Header[t.headerName][0] = randomUUID
	}

	t.next.ServeHTTP(rw, req)
}
