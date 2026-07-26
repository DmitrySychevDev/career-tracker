package main

import (
	"log"

	"github.com/DmitrySychevDev/career-tracker/internal/config"
	"github.com/DmitrySychevDev/career-tracker/internal/database"
	"github.com/DmitrySychevDev/career-tracker/internal/server"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	postgresDb, err := database.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := postgresDb.Close(); err != nil {
			log.Printf("failed to close database connection: %v", err)
		}
	}()

	log.Println("Connected to database")

	srv := server.New(*cfg, postgresDb.SQLDB)

	log.Printf("starting HTTP server on %s", srv.Addr)

	if err := server.RunServer(srv); err != nil {
		log.Fatal(err)
	}

}
