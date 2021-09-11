package rest

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mohammadne/go-pkgs/failures"
	"github.com/mohammadne/go-pkgs/logger"
)

func (rest *restEcho) authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token, failure := extractToken(ctx.Request())
		if failure != nil {
			rest.logger.Error("invalid token", logger.Error(failure))
			return ctx.JSON(failure.Status(), failure)
		}

		userId, failure := rest.authGrpc.GetTokenMetadata(token)
		if failure != nil || userId == 0 {
			return ctx.JSON(failure.Status(), failure)
		}

		ctx.Set("self_token", strconv.FormatUint(userId, 10))
		return next(ctx)
	}
}

var (
	missingAuth = failures.Rest{}.NewUnauthorized("authorization header is missing")
	invalidAuth = failures.Rest{}.NewUnauthorized("authorization header is malformed")
)

func extractToken(r *http.Request) (string, failures.Failure) {
	bearToken := r.Header.Get("Authorization")

	if bearToken == "" {
		return "", missingAuth
	}

	strArr := strings.Split(bearToken, " ")
	if len(strArr) != 2 {
		return "", invalidAuth
	}

	return strArr[1], nil
}
