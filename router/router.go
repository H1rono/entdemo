package router

import "github.com/labstack/echo/v4"

type Router struct{}

func New() *Router {
	return &Router{}
}

func (r *Router) SetupRoutes(e *echo.Echo) {
	e.GET("/ping", r.Ping)
}
