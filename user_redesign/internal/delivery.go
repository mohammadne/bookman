package internal

import "github.com/labstack/echo/v4"

type IUserHandler interface {
	Create() echo.HandlerFunc
	Get() echo.HandlerFunc
}
