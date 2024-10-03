package main

import (
	"context"
	"log"

	"github.com/H1rono/entdemo/repository"
	"github.com/H1rono/entdemo/router"
	"github.com/labstack/echo/v4"

	_ "github.com/lib/pq"
)

func main() {
	dbConfig := repository.ConfigFromEnv()
	if dbConfig == nil {
		log.Fatal("Missing required environment variables")
	}
	repo, err := repository.Connect(dbConfig)
	if err != nil {
		log.Fatalf("failed connecting to database: %v", err)
	}
	defer repo.Close()
	if err := repo.Migrate(context.Background()); err != nil {
		log.Fatalf("failed running migrations: %v", err)
	}
	e := echo.New()
	r := router.New(repo)
	r.SetupRoutes(e)
	log.Fatal(e.Start(":1323"))
}
