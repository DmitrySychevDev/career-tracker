package middleware

import (
	"log/slog"
	"net/http"

	"github.com/DmitrySychevDev/career-tracker/internal/httpx"
)

func RecoveryMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseWrapper := newResponseWriter(w)

		defer func() {
			err := recover()

			if err != nil {
				_ = httpx.WriteError(responseWrapper, http.StatusInternalServerError, "internal server error")

				requestId := RequestIDFromContext(r.Context())
				logger.Error("panic recovered.", "error", err, "method", r.Method,
					"path", r.URL.Path,
					"status", responseWrapper.StatusCode(),
					"request_id", requestId)
			}

		}()

		next.ServeHTTP(responseWrapper, r)
	})
}
