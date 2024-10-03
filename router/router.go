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
	e.GET("/ping", r.Ping)
}
