package middleware

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func RecoveryMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseWrapper := newResponseWriter(w)

		defer func() {
			err := recover()

			if err != nil {
				responseWrapper.WriteHeader(http.StatusInternalServerError)
				responseWrapper.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(responseWrapper).Encode(map[string]string{"message": "internal server error"})

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
