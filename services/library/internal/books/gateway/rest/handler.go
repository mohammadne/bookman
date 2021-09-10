package rest

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/library/internal/books"
	"github.com/mohammadne/bookman/library/pkg/failures"
)

type Handler interface {
	Route(group *echo.Group)

	get(ctx echo.Context) error
}

type handler struct {
	usecase books.Usecase
}

func NewHandler(usecase books.Usecase) Handler {
	return &handler{usecase: usecase}
}

func (h *handler) Route(group *echo.Group) {
	group.GET("/:id", h.get)
}

func (h *handler) get(ctx echo.Context) error {
	idStr := ctx.Param("id")
	if idStr == "" {
		failure := failures.Web{}.NewBadRequest("invalid id is given")
		return ctx.JSON(failure.Status(), failure)
	}

	id, parseErr := strconv.ParseUint(idStr, 10, 64)
	if parseErr != nil {
		failure := failures.Web{}.NewBadRequest("given id is malformed")
		return ctx.JSON(failure.Status(), failure)
	}

	book, failure := h.usecase.GetById(ctx.Request().Context(), id)
	if failure != nil {
		return ctx.JSON(failure.Status(), failure)
	}

	return ctx.JSON(http.StatusOK, book)
}
