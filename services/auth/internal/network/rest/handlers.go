package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mohammadne/go-pkgs/failures"
)

var (
	failureInvalidBody         = failures.Rest{}.NewBadRequest("invalid json body provided")
	failureInvalidCredentials  = failures.Rest{}.NewUnauthorized("invalid credentials given")
	failureUnprocessableEntity = failures.Rest{}.NewUnauthorized("unprocessable entity")
)

type User struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// TODO: REMOVE IT
var sampleUser = User{
	ID:       1,
	Email:    "email",
	Password: "password",
}

func (e echoRestAPI) signUp(ctx echo.Context) error {
	return nil
}

func (e echoRestAPI) signIn(ctx echo.Context) error {
	user := new(User)
	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(failureInvalidBody.Status(), failureInvalidBody)
	}

	// TODO: compare provided credentials with user-service and check there's a match
	if user.Email != sampleUser.Email || user.Password != sampleUser.Password {
		return ctx.JSON(failureInvalidCredentials.Status(), failureInvalidCredentials)
	}

	tokenDetail, err := e.jwt.CreateTokenDetail(user.ID)
	if err != nil {
		return ctx.JSON(failureUnprocessableEntity.Status(), failureUnprocessableEntity)
	}

	saveErr := e.cache.SetTokenDetail(user.ID, tokenDetail)
	if saveErr != nil {
		return ctx.JSON(failureUnprocessableEntity.Status(), failureUnprocessableEntity)
	}

	tokens := map[string]string{
		"access_token":  tokenDetail.AccessToken.Token,
		"refresh_token": tokenDetail.RefreshToken.Token,
	}

	return ctx.JSON(http.StatusOK, tokens)
}

func (e echoRestAPI) signOut(ctx echo.Context) error {
	return nil
}
