package home

import (
	"log"
	"net/http"
	"time"
)

const message = "Hello world!"

// Handlers struct
type Handlers struct {
	logger *log.Logger
}

// Handler - serves a message
func (h *Handlers) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

// Logger - prints request processing time to console
func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}

// SetupRoutes - declares the various active endpoints
func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.Handle("/", http.FileServer(http.Dir(".")))
	mux.HandleFunc("/hello", h.Logger(h.Handler))
}

// NewHandlers - inits a new set of handlers
func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
