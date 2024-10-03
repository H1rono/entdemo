package main

import (
	"context"
	"log"
	"time"

	"github.com/H1rono/entdemo/repository"

	_ "github.com/lib/pq"
)

func main() {
	dbConfig := repository.ConfigFromEnv()
	if dbConfig == nil {
		log.Fatal("missing required environment variables")
	}
	repo, err := repository.Connect(dbConfig)
	if err != nil {
		log.Fatalf("failed connecting to database: %v", err)
	}
	defer repo.Close()
	if err := repo.Migrate(context.Background()); err != nil {
		log.Fatalf("failed running migrations: %v", err)
	}
	res, err := repo.CreateCar(context.Background(), &repository.CreateCar{
		Model:        "Toyota",
		RegisteredAt: time.Now(),
	})
	if err != nil {
		log.Fatalf("failed creating car: %v", err)
	}
	log.Printf("car created: %+v\n", res)
}
