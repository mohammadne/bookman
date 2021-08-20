package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/user/internal"
	"github.com/mohammadne/bookman/user/internal/entities"
	"github.com/mohammadne/bookman/user/pkg/errors"
	"github.com/mohammadne/bookman/user/pkg/logger"
)

type userHandler struct {
	usecase internal.IUserUsecase
	logger  logger.Logger
}

// NewUserHandler
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

	result, err := h.usecase.Create(user)
	if err != nil {
		return ctx.JSON(err.Status, err)
	}

	return ctx.JSON(http.StatusCreated, result)
}

func (h *userHandler) Get(ctx echo.Context) error {
	userId, parseErr := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if parseErr != nil {
		restErr := errors.NewBadRequestError("user id sould be a number")
		return ctx.JSON(restErr.Status, restErr)
	}

	user, getErr := h.usecase.Get(userId)
	if getErr != nil {
		return ctx.JSON(getErr.Status, getErr)
	}

	return ctx.JSON(http.StatusCreated, user)
}
