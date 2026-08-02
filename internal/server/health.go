package server

import (
	"log/slog"
	"net/http"

	"github.com/DmitrySychevDev/career-tracker/internal/httpx"
)

func healthCheck(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		if err := httpx.WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"}); err != nil {
			logger.Error("failed to write health response", "error", err)
		}
	}
}
