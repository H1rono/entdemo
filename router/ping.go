package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GET /ping
func (r *Router) Ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
