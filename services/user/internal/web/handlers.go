package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/user/internal/models"
)

// Get is responsible to provide HTTP Get Location functionality
func (wh *echoWebHandler) get(ctx echo.Context) error {
	user := &models.User{
		Id:        1,
		FirstName: "mohammad",
		LastName:  "Nasr",
	}

	err := (*wh.database).CreateUser(user)
	if err != nil {
		(*wh.logger).Fatal("fatal")
	}

	return ctx.JSON(http.StatusOK, user)
}
