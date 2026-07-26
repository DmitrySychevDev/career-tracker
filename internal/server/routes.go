package server

import (
	"database/sql"
	"net/http"
)

func NewRouter(db *sql.DB) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthCheck)
	mux.Handle("GET /ready", readyHandler(db))

	return mux
}
