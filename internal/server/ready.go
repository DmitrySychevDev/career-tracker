package server

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/DmitrySychevDev/career-tracker/internal/httpx"
)

func readyHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		if err := db.PingContext(ctx); err != nil {
			_ = httpx.WriteJSON(w, http.StatusServiceUnavailable, map[string]string{"status": "not ready"})
			return
		}

		_ = httpx.WriteJSON(w, http.StatusOK, map[string]string{"status": "ready"})
	})
}
