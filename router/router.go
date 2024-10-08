package router

import (
	"github.com/H1rono/entdemo/repository"
	"github.com/labstack/echo/v4"
)

type Router struct {
	r *repository.Repository
}

func New(repo *repository.Repository) *Router {
	return &Router{repo}
}

func (r *Router) SetupRoutes(e *echo.Echo) {
	root := e.Group("/api")
	root.GET("/ping", r.Ping)
	{
		users := root.Group("/users")
		r.SetupUserRoutes(users)
	}
}
