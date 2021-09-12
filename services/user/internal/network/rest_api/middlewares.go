package rest_api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/user/pkg/failures"
	"github.com/mohammadne/bookman/user/pkg/logger"
)

func (rest *restEcho) authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	// span, ctx := rest.tra

	return func(ctx echo.Context) error {
		token, failure := extractToken(ctx.Request())
		if failure != nil {
			rest.logger.Error("invalid token", logger.Error(failure))
			return ctx.JSON(failure.Status(), failure)
		}
		// TODO : pass context
		userId, failure := rest.authGrpc.GetTokenMetadata(nil, token)
		if failure != nil || userId == 0 {
			return ctx.JSON(failure.Status(), failure)
		}

		ctx.Set("self_token", strconv.FormatUint(userId, 10))
		return next(ctx)
	}
}

var (
	missingAuth = failures.Network{}.NewUnauthorized("authorization header is missing")
	invalidAuth = failures.Network{}.NewUnauthorized("authorization header is malformed")
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
