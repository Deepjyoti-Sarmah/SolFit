package main

import (
	"log"

	"github.com/Deepjyoti-Sarmah/sol-kit-backend/api"
	"github.com/Deepjyoti-Sarmah/sol-kit-backend/config"
	db "github.com/Deepjyoti-Sarmah/sol-kit-backend/internal/models"
	"github.com/Deepjyoti-Sarmah/sol-kit-backend/storage"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// db connection
	dbPool, err := storage.NewPostgresDb(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbPool.Close()

	queries := db.New(dbPool)

	server := api.NewServer(cfg, dbPool, queries)
	if err := server.Start(); err != nil {
		return nil
	}

	return nil
}
