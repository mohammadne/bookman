package rest_api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/user/internal/models"
	"github.com/mohammadne/bookman/user/pkg/failures"
	"github.com/mohammadne/bookman/user/pkg/logger"
)

func (rest *restServer) getUser(c echo.Context) error {
	spanName := "network.rest_api.handlers.get_user"
	ctx, span := rest.tracer.Start(c.Request().Context(), spanName)
	defer span.End()

	user, failure := rest.getUserByIdString(ctx, c.Param("id"))
	if failure != nil {
		span.RecordError(failure)
		return c.JSON(failure.Status(), failure)
	}

	return c.JSON(http.StatusOK, user.Marshall(true))
}

func (rest *restServer) getMyUser(c echo.Context) error {
	spanName := "network.rest_api.handlers.get_my_user"
	ctx, span := rest.tracer.Start(c.Request().Context(), spanName)
	defer span.End()

	user, failure := rest.getUserByIdString(ctx, c.Get("self_token").(string))
	if failure != nil {
		rest.logger.Error(failure.Message(), logger.Error(failure))
		span.RecordError(failure)
		return c.JSON(failure.Status(), failure)
	}

	return c.JSON(http.StatusOK, user.Marshall(false))
}

func (rest *restServer) getUserByIdString(ctx context.Context, idStr string) (*models.User, failures.Failure) {
	if idStr == "" {
		return nil, failures.Network{}.NewBadRequest("invalid id is given")
	}

	id, parseErr := strconv.ParseUint(idStr, 10, 64)
	if parseErr != nil {
		return nil, failures.Network{}.NewBadRequest("given user id is malformed")
	}

	user, failure := rest.storage.FindUserById(ctx, id)
	if failure != nil {
		return nil, failure
	}

	return user, nil
}
