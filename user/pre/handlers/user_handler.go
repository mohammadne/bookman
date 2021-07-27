package handlers

import (
	"net/http"

	"github.com/go-delve/delve/pkg/proc/core"
	"github.com/labstack/echo"
	"github.com/mohammadne/bookman/user_redesign/domain/entities"
	// "github.com/mohammadne/bookman/user/domain/users"
	// "github.com/mohammadne/bookman/user/services"
	// "github.com/mohammadne/bookman/user/utils"
)

type UserHandler struct{}

// curl -X POST localhost:8080/users \
// -H 'Content-Type: application/json'\
// -d '{"id": 123}'
func (c *UserHandler) CreateUser(ctx echo.Context) error {

	user := new(entities.User)
	if err := ctx.Bind(user); err != nil {
		restErr := core.NewBadRequestError("invalid json body")
		return ctx.JSON(restErr.Status, restErr)
	}

	dao := users.MySQLUserDao{}
	result, saveErr := services.CreateUser(user, dao)
	if saveErr != nil {
		return ctx.JSON(saveErr.Status, saveErr)
	}

	return ctx.JSON(http.StatusCreated, result)
}

// curl -X GET localhost:8080/users
func (c *UserHandler) GetUser(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "GetUser")
}
