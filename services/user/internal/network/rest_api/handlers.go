package rest_api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/user/internal/models"
	"github.com/mohammadne/bookman/user/pkg/failures"
	"github.com/mohammadne/bookman/user/pkg/logger"
)

// get is responsible to provide HTTP Get Location functionality
func (rest *restEcho) getUser(ctx echo.Context) error {
	user, failure := rest.getUserByIdString(ctx.Param("id"))
	if failure != nil {
		return ctx.JSON(failure.Status(), failure)
	}

	return ctx.JSON(http.StatusOK, user.Marshall(true))
}

func (rest *restEcho) getMyUser(ctx echo.Context) error {
	user, failure := rest.getUserByIdString(ctx.Get("self_token").(string))
	if failure != nil {
		rest.logger.Error(failure.Message(), logger.Error(failure))
		return ctx.JSON(failure.Status(), failure)
	}

	return ctx.JSON(http.StatusOK, user.Marshall(false))
}

func (rest *restEcho) searchUsers(ctx echo.Context) error {
	return nil
}

func (rest *restEcho) getUserByIdString(idStr string) (*models.User, failures.Failure) {
	if idStr == "" {
		return nil, failures.Rest{}.NewBadRequest("invalid id is given")
	}

	id, parseErr := strconv.ParseInt(idStr, 10, 64)
	if parseErr != nil {
		return nil, failures.Rest{}.NewBadRequest("given user id is malformed")
	}

	user, failure := rest.database.FindUserById(id)
	if failure != nil {
		return nil, failure
	}

	return user, nil
}
