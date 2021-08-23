package rest

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get is responsible to provide HTTP Get Location functionality
func (rest *echoRestAPI) get(ctx echo.Context) error {
	idStr := ctx.Param("id")
	if idStr == "" {
		rest.logger.Error("user id is nil")
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	id, parseErr := strconv.ParseInt(idStr, 10, 64)
	if parseErr != nil {
		rest.logger.Error("user id is malformed")
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	user, readErr := rest.database.ReadUserById(id)
	if readErr != nil {
		return ctx.JSON(readErr.Status(), readErr)
	}

	return ctx.JSON(http.StatusOK, user)
}

func (rest *echoRestAPI) getMe(ctx echo.Context) error {
	return nil
}