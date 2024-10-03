package main

import (
	"context"
	"fmt"

	"github.com/H1rono/entdemo/repository"

	_ "github.com/lib/pq"
)

func main() {
	config := repository.ConfigFromEnv()
	if config == nil {
		panic("missing required environment variables")
	}
	repo, err := repository.Connect(config)
	if err != nil {
		panic(err)
	}
	defer repo.Close()
	if err := repo.Migrate(context.Background()); err != nil {
		panic(err)
	}
	user := &repository.CreateUser{
		Age:  20,
		Name: "H1rono",
	}
	u, err := repo.CreateUser(context.Background(), user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("user created: %+v\n", u)
}
