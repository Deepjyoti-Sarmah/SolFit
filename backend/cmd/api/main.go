package main

import (
	"log"

	"github.com/Deepjyoti-Sarmah/sol-kit-backend/config"
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
	dbConn, err := storage.NewPostgresDb(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	return nil
}
