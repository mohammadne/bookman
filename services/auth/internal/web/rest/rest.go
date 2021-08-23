package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/auth/internal/cache"
	"github.com/mohammadne/bookman/auth/internal/jwt"
	"github.com/mohammadne/go-pkgs/logger"
)

type Rest interface {
	SetupRoutes()
	Start()
}

type echoRestAPI struct {
	// injected parameters
	config *Config
	logger logger.Logger
	cache  cache.Cache
	jwt    jwt.Jwt

	// internal dependencies
	instance *echo.Echo
}

func NewEcho(cfg *Config, log logger.Logger, c cache.Cache, j jwt.Jwt) Rest {
	handler := &echoRestAPI{config: cfg, logger: log, cache: c, jwt: j}

	handler.instance = echo.New()
	handler.instance.HideBanner = true

	return handler
}

func (rest *echoRestAPI) SetupRoutes() {
	rest.instance.POST("/auth/login", rest.login)
}

func (rest *echoRestAPI) Start() {
	rest.logger.Info(
		"starting server",
		logger.String("address", rest.config.URL),
	)

	go func() {
		if err := rest.instance.Start(rest.config.URL); err != nil {
			rest.logger.Fatal("starting server failed", logger.Error(err))
		}
	}()
}
