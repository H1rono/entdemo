package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/H1rono/entdemo/repository"
	_ "github.com/lib/pq"
)

const (
	CMD_CREATE  = "create"
	CMD_GET_ALL = "get-all"
	CMD_GET     = "get"
	CMD_DELETE  = "delete"
)

// Command seal
type Command interface {
	commandTag()
}

type CreateCommand struct {
	Age  int
	Name string
}

func (*CreateCommand) commandTag() {}

type GetAllCommand struct{}

func (*GetAllCommand) commandTag() {}

type GetCommand struct {
	ID int
}

func (*GetCommand) commandTag() {}

type DeleteCommand struct {
	ID int
}

func (*DeleteCommand) commandTag() {}

func CreateCommandFlagSet(o *CreateCommand) *flag.FlagSet {
	fs := flag.NewFlagSet(CMD_CREATE, flag.ExitOnError)
	fs.IntVar(&o.Age, "age", 0, "age")
	fs.StringVar(&o.Name, "name", "", "name")
	return fs
}

func GetAllCommandFlagSet(_ *GetAllCommand) *flag.FlagSet {
	return flag.NewFlagSet(CMD_GET_ALL, flag.ExitOnError)
}

func GetCommandFlagSet(o *GetCommand) *flag.FlagSet {
	fs := flag.NewFlagSet(CMD_GET, flag.ExitOnError)
	fs.IntVar(&o.ID, "id", 0, "id")
	return fs
}

func DeleteCommandFlagSet(o *DeleteCommand) *flag.FlagSet {
	fs := flag.NewFlagSet(CMD_DELETE, flag.ExitOnError)
	fs.IntVar(&o.ID, "id", 0, "id")
	return fs
}

func Parse() (Command, error) {
	if len(os.Args) < 2 {
		return nil, fmt.Errorf("missing command")
	}
	switch os.Args[1] {
	case CMD_CREATE:
		cmd := CreateCommand{}
		fs := CreateCommandFlagSet(&cmd)
		if err := fs.Parse(os.Args[2:]); err != nil {
			return nil, err
		}
		return &cmd, nil
	case CMD_GET_ALL:
		cmd := GetAllCommand{}
		fs := GetAllCommandFlagSet(&cmd)
		if err := fs.Parse(os.Args[2:]); err != nil {
			return nil, err
		}
		return &cmd, nil
	case CMD_GET:
		cmd := GetCommand{}
		fs := GetCommandFlagSet(&cmd)
		if err := fs.Parse(os.Args[2:]); err != nil {
			return nil, err
		}
		return &cmd, nil
	case CMD_DELETE:
		cmd := DeleteCommand{}
		fs := DeleteCommandFlagSet(&cmd)
		if err := fs.Parse(os.Args[2:]); err != nil {
			return nil, err
		}
		return &cmd, nil
	default:
		return nil, fmt.Errorf("unknown command: %s", flag.Args()[0])
	}
}

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

	cmd, err := Parse()
	if err != nil {
		log.Fatalf("failed parsing command: %v", err)
	}
	switch cmd.(type) {
	case *CreateCommand:
		log.Printf("create command: %+v", cmd)
		c := cmd.(*CreateCommand)
		u, err := repo.CreateUser(context.Background(), &repository.CreateUser{
			Age:  c.Age,
			Name: c.Name,
		})
		if err != nil {
			log.Fatalf("failed creating user: %v", err)
		}
		log.Printf("created user: %+v", u)
	case *GetAllCommand:
		log.Printf("get-all command: %+v", cmd)
		users, err := repo.GetUsers(context.Background())
		if err != nil {
			log.Fatalf("failed getting users: %v", err)
		}
		log.Println("got users:")
		for _, u := range users {
			log.Printf("\t%+v", u)
		}
	case *GetCommand:
		log.Printf("get command: %+v", cmd)
		c := cmd.(*GetCommand)
		u, err := repo.GetUser(context.Background(), c.ID)
		if err != nil {
			log.Fatalf("failed getting user: %v", err)
		}
		log.Printf("got user: %+v", u)
	case *DeleteCommand:
		log.Printf("delete command: %+v", cmd)
		c := cmd.(*DeleteCommand)
		err := repo.DeleteUser(context.Background(), c.ID)
		if err != nil {
			log.Fatalf("failed deleting user: %v", err)
		}
		log.Printf("deleted user: %+d", c.ID)
	}
}
