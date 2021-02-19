package handlers

import (
	"compress/gzip"
	"net/http"
	"strings"
)

// GzipHandler ...
type GzipHandler struct {
}

// GzipMiddleware ...
func (g *GzipHandler) GzipMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			wrw := NewWrappedResponseWriter(rw)
			wrw.Header().Set("Content-Encoding", "gzip")

			next.ServeHTTP(wrw, r)
			defer wrw.Flush()
			return
		}

		next.ServeHTTP(rw, r)
	})
}

// WrappedResponseWriter ...
type WrappedResponseWriter struct {
	rw http.ResponseWriter
	gw *gzip.Writer
}

// NewWrappedResponseWriter ...
func NewWrappedResponseWriter(rw http.ResponseWriter) *WrappedResponseWriter {
	gw := gzip.NewWriter(rw)

	return &WrappedResponseWriter{rw: rw, gw: gw}
}

// Header ...
func (wr *WrappedResponseWriter) Header() http.Header {
	return wr.rw.Header()
}

// Write ...
func (wr *WrappedResponseWriter) Write(d []byte) (int, error) {
	return wr.gw.Write(d)
}

// WriteHeader ...
func (wr *WrappedResponseWriter) WriteHeader(statuscode int) {
	wr.rw.WriteHeader(statuscode)
}

// Flush ...
func (wr *WrappedResponseWriter) Flush() {
	wr.gw.Flush()
	wr.gw.Close()
}
