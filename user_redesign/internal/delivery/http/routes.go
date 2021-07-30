package http

import (
	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/user/internal"
)

// Map comments routes
// h comments.Handlers, mw *middleware.MiddlewareManager
func MapUserRoutes(group *echo.Group, h internal.IUserHandler) {
	group.POST("/users", h.Create)
	group.GET("/users/:user_id", h.Get)
}
