package rest_api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/library/pkg/failures"
	"github.com/mohammadne/bookman/library/pkg/logger"
)

func (rest *restApi) authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		spanName := "network.rest_api.middlewares.authenticate"
		ctx, span := rest.tracer.Start(c.Request().Context(), spanName)
		defer span.End()

		token, failure := extractToken(c.Request())
		if failure != nil {
			rest.logger.Error("invalid token", logger.Error(failure))
			span.RecordError(failure)
			return c.JSON(failure.Status(), failure)
		}

		userId, failure := rest.authGrpc.GetTokenMetadata(ctx, token)
		if failure != nil || userId == 0 {
			span.RecordError(failure)
			return c.JSON(failure.Status(), failure)
		}

		c.Set("self_token", strconv.FormatUint(userId, 10))
		return next(c)
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
