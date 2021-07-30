package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/user/internal"
	"github.com/mohammadne/bookman/user/internal/entities"
	"github.com/mohammadne/bookman/user/pkg/errors"
)

type userHandler struct {
	usecase internal.IUserUsecase
	// logger logger.Logger
}

//
func NewUserHandler(usecase internal.IUserUsecase) *userHandler {
	return &userHandler{
		usecase: usecase,
	}
}

func (h *userHandler) Create(ctx echo.Context) error {
	user := new(entities.User)
	if err := ctx.Bind(user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		return ctx.JSON(restErr.Status, restErr)
	}

	result, err := h.usecase.Create(nil, user)
	if err != nil {
		return ctx.JSON(err.Status, err)
	}

	return ctx.JSON(http.StatusCreated, result)
}

func (h *userHandler) Get(ctx echo.Context) error {
	return nil
}
