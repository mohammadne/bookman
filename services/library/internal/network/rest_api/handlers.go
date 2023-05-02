package rest_api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/library/pkg/failures"
	"github.com/mohammadne/bookman/library/pkg/logger"
)

func getSpanName(method string) string {
	const spanName = "network.rest_api.handlers"
	return spanName + "." + method
}

func (rest *restApi) getBook(c echo.Context) error {
	ctx, span := rest.tracer.Start(c.Request().Context(), getSpanName("get_user"))
	defer span.End()

	id, failure := idFromParams(c)
	if failure != nil {
		rest.logger.Error(failure.Message(), logger.Error(failure))
		span.RecordError(failure)
		return c.JSON(failure.Status(), failure)
	}

	book, failure := rest.database.GetBook(ctx, id)
	if failure != nil {
		rest.logger.Error(failure.Message(), logger.Error(failure))
		span.RecordError(failure)
		return c.JSON(failure.Status(), failure)
	}

	return c.JSON(http.StatusOK, book)
}

func (rest *restApi) getAuthor(c echo.Context) error {
	ctx, span := rest.tracer.Start(c.Request().Context(), getSpanName("get_author"))
	defer span.End()

	id, failure := idFromParams(c)
	if failure != nil {
		rest.logger.Error(failure.Message(), logger.Error(failure))
		span.RecordError(failure)
		return c.JSON(failure.Status(), failure)
	}

	book, failure := rest.database.GetBook(ctx, id)
	if failure != nil {
		rest.logger.Error(failure.Message(), logger.Error(failure))
		span.RecordError(failure)
		span.RecordError(failure)
		return c.JSON(failure.Status(), failure)
	}

	return c.JSON(http.StatusOK, book)
}

func idFromParams(c echo.Context) (uint64, failures.Failure) {
	idStr := c.Param("id")
	if idStr == "" {
		return 0, failures.Network{}.NewBadRequest("invalid id is given")
	}

	id, parseErr := strconv.ParseUint(idStr, 10, 64)
	if parseErr != nil {
		return 0, failures.Network{}.NewBadRequest("given user id is malformed")
	}

	return id, nil
}
