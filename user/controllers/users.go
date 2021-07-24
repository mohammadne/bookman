package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/mohammadne/bookman/user/domain/users"
	"github.com/mohammadne/bookman/user/services"
)

// curl -X POST localhost:8080/users -d '{"id": 123}'
func (c *Controller) CreateUser(ctx echo.Context) error {
	user := new(users.User)
	if err := ctx.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// bytes, err := ioutil.ReadAll(ctx.Request().Body)
	// if err != nil {
	// 	return err
	// }

	// var user users.User
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	return err
	// }

	result, saveErr := services.CreateUser(*user)
	if saveErr != nil {
		return saveErr
	}

	return ctx.JSON(http.StatusCreated, result)
}

// curl -X GET localhost:8080/users
func (c *Controller) GetUser(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "GetUser")
}
