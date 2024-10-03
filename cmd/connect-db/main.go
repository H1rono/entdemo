package main

import (
	"context"
	"log"

	repo "github.com/H1rono/entdemo/repository"

	_ "github.com/lib/pq"
)

func main() {
	dbConfig := repo.ConfigFromEnv()
	if dbConfig == nil {
		log.Fatal("Missing required environment variables")
	}
	repository, err := repo.Connect(dbConfig)
	if err != nil {
		log.Fatalf("failed connecting to database: %v", err)
	}
	defer func() {
		if err := repository.Close(); err != nil {
			log.Fatalf("failed closing connection: %v", err)
		}
	}()
	if err := repository.Migrate(context.Background()); err != nil {
		log.Fatalf("failed migration: %v", err)
	}
}
