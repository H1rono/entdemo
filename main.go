package main

import (
	"github.com/H1rono/entdemo/router"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	e := echo.New()
	r := router.New()
	r.SetupRoutes(e)
	log.Fatal(e.Start(":1323"))
}
