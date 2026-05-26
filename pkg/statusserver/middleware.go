package statusserver

import (
	"net/http"

	"github.com/go-logr/logr"
)

type Middleware func(http.Handler) http.Handler

// chain applies middleware in order: first middleware wraps second, etc.
func chain(h http.Handler, middlewares ...Middleware) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// recoveryMiddleware recovers from panics in HTTP handlers to prevent Server crashes.
func recoveryMiddleware(log logr.Logger) Middleware {
	_ = "STUB: not implemented"
	return *new(Middleware)
}

// loggingMiddleware logs incoming HTTP requests.
func loggingMiddleware(log logr.Logger) Middleware {
	_ = "STUB: not implemented"
	return *new(Middleware)
}

// bodySizeLimitMiddleware enforces a maximum request body size.
func bodySizeLimitMiddleware(log logr.Logger, maxBytes int64) Middleware {
	_ = "STUB: not implemented"
	return *new(Middleware)
}

// Reject based on Content-Length header if present

// Wrap body to enforce limit for chunked/streaming requests
