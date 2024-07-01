package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/H1rono/entdemo/ent"

	_ "github.com/lib/pq"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Db       string
}

func DbConfigFromEnv() *DbConfig {
	host, ok := os.LookupEnv("POSTGRES_HOST")
	if !ok {
		host = "localhost"
	}
	port, ok := os.LookupEnv("POSTGRES_PORT")
	if !ok {
		port = "5432"
	}
	user, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		return nil
	}
	password, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		return nil
	}
	db, ok := os.LookupEnv("POSTGRES_DB")
	if !ok {
		return nil
	}
	return &DbConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Db:       db,
	}
}

func (c *DbConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", c.Host, c.Port, c.User, c.Db, c.Password)
}

func main() {
	dbConfig := DbConfigFromEnv()
	if dbConfig == nil {
		log.Fatal("Missing required environment variables")
	}
	dbClient, err := ent.Open("postgres", dbConfig.DSN())
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer func() {
		if err := dbClient.Close(); err != nil {
			log.Fatalf("failed closing connection to postgres: %v", err)
		}
	}()
	if err := dbClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
