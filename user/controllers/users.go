package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func (c *Controller) CreateUser(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "pong")
}

func (c *Controller) GetUser(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "pong")
}
