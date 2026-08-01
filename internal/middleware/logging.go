package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func LoggingMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		requestId := RequestIDFromContext(r.Context())

		logger.Info("request started",
			"method", r.Method,
			"path", r.URL.Path,
			"request_id", requestId,
		)

		responseWrapper := newResponseWriter(w)
		next.ServeHTTP(responseWrapper, r)

		logger.Info("request completed",
			"method", r.Method,
			"path", r.URL.Path,
			"duration", time.Since(start).String(),
			"status", responseWrapper.StatusCode(),
			"request_id", requestId,
		)
	})
}
