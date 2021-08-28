package rest

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/auth/internal/jwt"
	"github.com/mohammadne/bookman/auth/internal/models"
	"github.com/mohammadne/go-pkgs/failures"
)

var (
	failureInvalidBody         = failures.Rest{}.NewBadRequest("invalid json body provided")
	failureBadRequest          = failures.Rest{}.NewBadRequest("email is already registered")
	failureNotFound            = failures.Rest{}.NewNotFound("invalid email and password credentials given")
	failureUnautorized         = failures.Rest{}.NewUnauthorized("unauthorized")
	failureUnprocessableEntity = failures.Rest{}.NewUnprocessableEntity("unprocessable entity")
)

func (r restServer) signUp(ctx echo.Context) error {
	userCredential := new(models.UserCredential)
	if err := ctx.Bind(userCredential); err != nil {
		return ctx.JSON(failureInvalidBody.Status(), failureInvalidBody)
	}

	userId, err := r.userGrpc.CreateUser(userCredential)
	if err != nil || userId == 0 {
		return ctx.JSON(failureBadRequest.Status(), failureBadRequest)
	}

	tokens, failure := r.generateTokens(userId)
	if failure != nil {
		return ctx.JSON(failure.Status(), failure)
	}

	return ctx.JSON(http.StatusOK, tokens)
}

func (r restServer) signIn(ctx echo.Context) error {
	userCredential := new(models.UserCredential)
	if err := ctx.Bind(userCredential); err != nil {
		return ctx.JSON(failureInvalidBody.Status(), failureInvalidBody)
	}

	userId, err := r.userGrpc.GetUser(userCredential)
	if err != nil || userId == 0 {
		return ctx.JSON(failureNotFound.Status(), failureNotFound)
	}

	tokens, failure := r.generateTokens(userId)
	if failure != nil {
		return ctx.JSON(failure.Status(), failure)
	}

	return ctx.JSON(http.StatusOK, tokens)
}

func (r restServer) signOut(ctx echo.Context) error {
	tokenString := extractToken(ctx.Request())
	accessDetails, err := r.jwt.ExtractTokenMetadata(tokenString, jwt.Access)
	if err != nil {
		return ctx.JSON(failureUnautorized.Status(), failureUnautorized)
	}

	deleted, delErr := r.cache.RevokeJwt(accessDetails.TokenUuid)
	if delErr != nil || deleted == 0 {
		return ctx.JSON(failureUnautorized.Status(), failureUnautorized)
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
