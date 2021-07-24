package main

import (
	"github.com/labstack/echo"
	"github.com/mohammadne/bookman/user/controllers"
)

func routeUrls(e *echo.Echo) {
	ctrl := controllers.Controller{}

	e.GET("/ping", ctrl.Ping)

	e.GET("/users/:user_id", ctrl.GetUser)
	e.POST("/users", ctrl.CreateUser)
}
