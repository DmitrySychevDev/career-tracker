package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/DmitrySychevDev/career-tracker/internal/config"
	"github.com/DmitrySychevDev/career-tracker/internal/database"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(map[string]string{"status": "ok"}); err != nil {
		log.Fatal(err)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", healthHandler)

	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	if err := sqlDb.Ping(); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := sqlDb.Close(); err != nil {
			log.Printf("failed to close database connection: %v", err)
		}
	}()

	log.Println("Connected to database")

	server := &http.Server{
		Addr:              ":" + cfg.AppPort,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("HTTP server started on %s", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
