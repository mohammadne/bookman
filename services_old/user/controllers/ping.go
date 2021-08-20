package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

// Ping is a test controller
// curl -X GET localhost:8080/ping
func (c *Controller) Ping(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "pong")
}
