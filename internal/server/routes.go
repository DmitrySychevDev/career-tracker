package server

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/DmitrySychevDev/career-tracker/internal/middleware"
)

func NewRouter(db *sql.DB, logger *slog.Logger) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthCheck(logger))
	mux.Handle("GET /ready", readyHandler(db))

	handler := middleware.LoggingMiddleware(logger, mux)
	handler = middleware.RecoveryMiddleware(logger, handler)
	handler = middleware.RequestIDMiddleware(handler)

	return handler
}
