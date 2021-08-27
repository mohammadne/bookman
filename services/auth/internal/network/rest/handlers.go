package rest

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/auth/internal/jwt"
	"github.com/mohammadne/go-pkgs/failures"
)

var (
	failureInvalidBody         = failures.Rest{}.NewBadRequest("invalid json body provided")
	failureInvalidCredentials  = failures.Rest{}.NewBadRequest("invalid credentials given")
	failureUnautorized         = failures.Rest{}.NewUnauthorized("unauthorized")
	failureUnprocessableEntity = failures.Rest{}.NewUnprocessableEntity("unprocessable entity")
)

type User struct {
	Id       uint64 `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// TODO: REMOVE IT
var sampleUser = User{
	Id:       1,
	Email:    "email",
	Password: "password",
}

func (e restServer) signUp(ctx echo.Context) error {
	return nil
}

func (e restServer) signIn(ctx echo.Context) error {
	user := new(User)
	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(failureInvalidBody.Status(), failureInvalidBody)
	}

	// TODO: compare provided credentials with user-service and check there's a match
	if user.Email != sampleUser.Email || user.Password != sampleUser.Password {
		return ctx.JSON(failureInvalidCredentials.Status(), failureInvalidCredentials)
	}

	tokens, failure := e.generateTokens(user.Id)
	if failure != nil {
		return ctx.JSON(failure.Status(), failure)
	}

	return ctx.JSON(http.StatusOK, tokens)
}

func (r restServer) signOut(ctx echo.Context) error {
	tokenString := extractToken(ctx.Request())
	accessDetails, err := r.jwt.ExtractTokenMetadata(tokenString, jwt.Access)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, "unauthorized")
	}

	deleted, delErr := r.cache.RevokeJwt(accessDetails.TokenUuid)
	if delErr != nil || deleted == 0 {
		return ctx.JSON(http.StatusUnauthorized, "unauthorized")
	}

	return ctx.JSON(http.StatusOK, "Successfully logged out")
}

func (r restServer) refreshToken(ctx echo.Context) error {
	mapToken := map[string]string{}
	if err := ctx.Bind(mapToken); err != nil {
		return ctx.JSON(failureInvalidBody.Status(), failureInvalidBody)
	}

	refreshToken, ok := mapToken["refresh_token"]
	if !ok {
		return ctx.JSON(failureInvalidBody.Status(), failureInvalidBody)
	}

	accessDetails, err := r.jwt.ExtractTokenMetadata(refreshToken, jwt.Refresh)
	if err != nil {
		return ctx.JSON(failureUnprocessableEntity.Status(), failureUnprocessableEntity)
	}

	userId, err := r.cache.RevokeJwt(accessDetails.TokenUuid)
	if err != nil || userId == 0 {
		return ctx.JSON(failureUnautorized.Status(), failureUnautorized)
	}

	tokens, failure := r.generateTokens(accessDetails.UserId)
	if failure != nil {
		return ctx.JSON(failure.Status(), failure)
	}

	return ctx.JSON(http.StatusOK, tokens)
}

func (r restServer) generateTokens(userId uint64) (map[string]string, failures.Failure) {
	jwt, err := r.jwt.CreateJwt(userId)
	if err != nil {
		return nil, failureUnprocessableEntity
	}

	saveErr := r.cache.SetJwt(userId, jwt)
	if saveErr != nil {
		return nil, failureUnprocessableEntity
	}

	return map[string]string{
		"access_token":  jwt.AccessToken.Token,
		"refresh_token": jwt.RefreshToken.Token,
	}, nil
}

// extractToken takes Authorization from the_token_xxx
func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}
