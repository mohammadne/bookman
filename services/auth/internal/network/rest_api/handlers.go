package rest_api

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/auth/internal/jwt"
	"github.com/mohammadne/bookman/auth/internal/models"
	"github.com/mohammadne/bookman/auth/pkg/failures"
)

var (
	failureInvalidBody         = failures.Network{}.NewBadRequest("invalid json body provided")
	failureBadRequest          = failures.Network{}.NewBadRequest("email is already registered")
	failureNotFound            = failures.Network{}.NewNotFound("invalid email and password credentials given")
	failureUnautorized         = failures.Network{}.NewUnauthorized("unauthorized")
	failureUnprocessableEntity = failures.Network{}.NewUnprocessableEntity("unprocessable entity")
)

func (handler restServer) signUp(c echo.Context) error {
	spanName := "network.rest_api.handlers.sign_up"
	ctx, span := handler.tracer.Start(c.Request().Context(), spanName)
	defer span.End()

	userCredential := new(models.UserCredential)
	if err := c.Bind(userCredential); err != nil {
		span.RecordError(failureInvalidBody)
		return c.JSON(failureInvalidBody.Status(), failureInvalidBody)
	}

	userId, failure := handler.userGrpc.CreateUser(ctx, userCredential)
	if failure != nil || userId == 0 {
		span.RecordError(failure)
		return c.JSON(failure.Status(), failure)
	}

	tokens, failure := handler.generateTokens(ctx, userId)
	if failure != nil {
		span.RecordError(failure)
		return c.JSON(failure.Status(), failure)
	}

	return c.JSON(http.StatusOK, tokens)
}

func (handler restServer) signIn(c echo.Context) error {
	spanName := "network.rest_api.handlers.sign_in"
	ctx, span := handler.tracer.Start(c.Request().Context(), spanName)
	defer span.End()

	userCredential := new(models.UserCredential)
	if err := c.Bind(userCredential); err != nil {
		span.RecordError(failureInvalidBody)
		return c.JSON(failureInvalidBody.Status(), failureInvalidBody)
	}

	userId, failure := handler.userGrpc.GetUser(ctx, userCredential)
	if failure != nil || userId == 0 {
		span.RecordError(failure)
		return c.JSON(failure.Status(), failure)
	}

	tokens, failure := handler.generateTokens(ctx, userId)
	if failure != nil {
		span.RecordError(failure)
		return c.JSON(failure.Status(), failure)
	}

	return c.JSON(http.StatusOK, tokens)
}

func (handler restServer) signOut(c echo.Context) error {
	spanName := "network.rest_api.handlers.sign_out"
	ctx, span := handler.tracer.Start(c.Request().Context(), spanName)
	defer span.End()

	tokenString := extractToken(c.Request())
	accessDetails, failure := handler.jwt.ExtractTokenMetadata(ctx, tokenString, jwt.Access)
	if failure != nil {
		span.RecordError(failure)
		return c.JSON(failure.Status(), failure)
	}

	deleted, failure := handler.cache.RevokeJwt(ctx, accessDetails.TokenUuid)
	if failure != nil || deleted == 0 {
		span.RecordError(failure)
		return c.JSON(failure.Status(), failure)
	}

	return c.JSON(http.StatusOK, "Successfully logged out")
}

func (handler restServer) refreshToken(c echo.Context) error {
	spanName := "network.rest_api.handlers.refresh_token"
	ctx, span := handler.tracer.Start(c.Request().Context(), spanName)
	defer span.End()

	body := struct {
		RefreshToken string `json:"refresh_token"`
	}{}

	if err := c.Bind(&body); err != nil {
		span.RecordError(err)
		return c.JSON(failureInvalidBody.Status(), failureInvalidBody)
	}

	accessDetails, failure := handler.jwt.ExtractTokenMetadata(ctx, body.RefreshToken, jwt.Refresh)
	if failure != nil {
		span.RecordError(failure)
		return c.JSON(failure.Status(), failure)
	}

	userId, failure := handler.cache.RevokeJwt(ctx, accessDetails.TokenUuid)
	if failure != nil || userId == 0 {
		span.RecordError(failure)
		return c.JSON(failure.Status(), failure)
	}

	tokens, failure := handler.generateTokens(ctx, accessDetails.UserId)
	if failure != nil {
		span.RecordError(failure)
		return c.JSON(failure.Status(), failure)
	}

	return c.JSON(http.StatusOK, tokens)
}

func (handler restServer) generateTokens(ctx context.Context, userId uint64) (map[string]string, failures.Failure) {
	jwt, failure := handler.jwt.CreateJwt(ctx, userId)
	if failure != nil {
		return nil, failure
	}

	failure = handler.cache.SetJwt(ctx, userId, jwt)
	if failure != nil {
		return nil, failure
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
