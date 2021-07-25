package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mohammadne/bookman/user/domain"
	"github.com/mohammadne/bookman/user/services"
	"github.com/mohammadne/bookman/user/utils"
)

// curl -X POST localhost:8080/users \
// -H 'Content-Type: application/json'\
// -d '{"id": 123}'
func (c *Controller) CreateUser(ctx echo.Context) error {
	user := new(domain.User)
	if err := ctx.Bind(user); err != nil {
		restErr := utils.NewBadRequest("invalid json body")
		return ctx.JSON(restErr.Status, restErr)
	}

	result, saveErr := services.CreateUser(*user)
	if saveErr != nil {
		return ctx.JSON(saveErr.Status, saveErr)
	}

	return ctx.JSON(http.StatusCreated, result)
}

// curl -X GET localhost:8080/users
func (c *Controller) GetUser(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "GetUser")
}
