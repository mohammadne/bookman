package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/auth/internal/cache"
	"github.com/mohammadne/bookman/auth/internal/jwt"
	"github.com/mohammadne/bookman/auth/internal/network"
	"github.com/mohammadne/go-pkgs/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type restServer struct {
	// injected parameters
	config *Config
	logger logger.Logger
	cache  cache.Cache
	jwt    jwt.Jwt

	// internal dependencies
	instance *echo.Echo
}

func New(cfg *Config, log logger.Logger, c cache.Cache, j jwt.Jwt) network.Server {
	handler := &restServer{config: cfg, logger: log, cache: c, jwt: j}

	handler.instance = echo.New()
	handler.instance.HideBanner = true
	handler.setupRoutes()

	return handler
}

func (rest *restServer) setupRoutes() {
	authGroup := rest.instance.Group("/auth")

	authGroup.POST("/metrics", echo.WrapHandler(promhttp.Handler()))
	authGroup.POST("/sign_up", rest.signUp)
	authGroup.POST("/sign_in", rest.signIn)
	authGroup.POST("/sign_out", rest.signOut)
}

func (rest *restServer) Serve(<-chan struct{}) {
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
