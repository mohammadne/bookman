package http

import (
	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/user/internal"
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

func (h *userHandler) Create() echo.HandlerFunc {
	return nil
}

func (h *userHandler) Get() echo.HandlerFunc {
	return nil
}
