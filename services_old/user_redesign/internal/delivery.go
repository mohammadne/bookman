package internal

import "github.com/labstack/echo/v4"

type IUserHandler interface {
	Create(ctx echo.Context) error
	Get(ctx echo.Context) error
}
