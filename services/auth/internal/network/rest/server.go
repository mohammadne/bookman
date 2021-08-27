package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/auth/internal/cache"
	"github.com/mohammadne/bookman/auth/internal/jwt"
	"github.com/mohammadne/bookman/auth/internal/network"
	grpc_client "github.com/mohammadne/bookman/auth/internal/network/grpc/clients"
	"github.com/mohammadne/go-pkgs/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type restServer struct {
	// injected parameters
	config   *Config
	logger   logger.Logger
	cache    cache.Cache
	jwt      jwt.Jwt
	userGrpc grpc_client.User

	// internal dependencies
	instance *echo.Echo
}

func New(cfg *Config, log logger.Logger, c cache.Cache, j jwt.Jwt, ug grpc_client.User) network.Server {
	handler := &restServer{config: cfg, logger: log, cache: c, jwt: j, userGrpc: ug}

	handler.instance = echo.New()
	handler.instance.HideBanner = true
	handler.setupRoutes()

	return handler
}

func (rest *restServer) setupRoutes() {
	rest.instance.POST("/auth/metrics", echo.WrapHandler(promhttp.Handler()))
	rest.instance.POST("/auth/sign_up", rest.signUp)
	rest.instance.POST("/auth/sign_in", rest.signIn)
	rest.instance.POST("/auth/sign_out", rest.signOut)
	rest.instance.POST("/auth/refresh_token", rest.refreshToken)
}

func (rest *restServer) Serve(<-chan struct{}) {
	rest.logger.Info(
		"starting server",
		logger.String("address", rest.config.URL),
	)

	if err := rest.instance.Start(rest.config.URL); err != nil {
		rest.logger.Fatal("starting server failed", logger.Error(err))
	}
}
