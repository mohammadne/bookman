package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Get is responsible to provide HTTP Get Location functionality
func (wh *echoWebHandler) get(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "worked")
}
