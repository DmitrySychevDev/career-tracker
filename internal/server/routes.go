package server

import (
	"database/sql"
	"log/slog"
	"net/http"
)

func NewRouter(db *sql.DB, logger *slog.Logger) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthCheck(logger))
	mux.Handle("GET /ready", readyHandler(db))

	return mux
}
