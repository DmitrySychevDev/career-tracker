package main

import (
	"log/slog"
	"os"

	"github.com/DmitrySychevDev/career-tracker/internal/config"
	"github.com/DmitrySychevDev/career-tracker/internal/database"
	"github.com/DmitrySychevDev/career-tracker/internal/server"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	postgresDb, err := database.NewPostgres(cfg)
	if err != nil {
		logger.Error("failed to connect to the db", "error", err)
		os.Exit(1)
	}

	defer func() {
		if err := postgresDb.Close(); err != nil {
			logger.Error("failed to close database connection", "error", err)
		}
	}()

	logger.Info("Connected to database")

	srv := server.New(*cfg, postgresDb.SQLDB, logger)

	logger.Info("starting HTTP server", "addr", srv.Addr)

	if err := server.RunServer(srv, logger); err != nil {
		logger.Error("unable to start the server", "error", err)
		os.Exit(1)
	}

}
