package repository

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/H1rono/entdemo/ent"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Db       string
}

func ConfigFromEnv() *DbConfig {
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

type Repository struct {
	c *ent.Client
}

func New(client *ent.Client) *Repository {
	return &Repository{client}
}

func Connect(c *DbConfig) (*Repository, error) {
	client, err := ent.Open("postgres", c.DSN())
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	return New(client), nil
}

func (r *Repository) Close() error {
	return r.c.Close()
}

func (r *Repository) Migrate(c context.Context) error {
	return r.c.Schema.Create(c)
}
