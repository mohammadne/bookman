package rest_api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/auth/internal/cache"
	"github.com/mohammadne/bookman/auth/internal/jwt"
	"github.com/mohammadne/bookman/auth/internal/network"
	"github.com/mohammadne/bookman/auth/internal/network/grpc"
	"github.com/mohammadne/bookman/auth/pkg/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/trace"
)

type restServer struct {
	config   *Config
	logger   logger.Logger
	tracer   trace.Tracer
	cache    cache.Cache
	jwt      jwt.Jwt
	userGrpc grpc.UserClient

	// internal dependencies
	echo *echo.Echo
}

func New(cfg *Config, log logger.Logger, t trace.Tracer, c cache.Cache, j jwt.Jwt, uc grpc.UserClient) network.Server {
	handler := &restServer{config: cfg, logger: log, tracer: t, cache: c, jwt: j, userGrpc: uc}

	handler.echo = echo.New()
	handler.echo.HideBanner = true
	handler.setupRoutes()

	return handler
}

func (rest *restServer) setupRoutes() {
	rest.echo.POST("/metrics", echo.WrapHandler(promhttp.Handler()))

	authGroup := rest.echo.Group("/auth")
	authGroup.POST("/auth/sign_up", rest.signUp)
	authGroup.POST("/auth/sign_in", rest.signIn)
	authGroup.POST("/auth/sign_out", rest.signOut)
	authGroup.POST("/auth/refresh_token", rest.refreshToken)
}

func (rest *restServer) Serve(<-chan struct{}) {
	Address := fmt.Sprintf("%s:%s", rest.config.Host, rest.config.Port)
	rest.logger.Info("starting server", logger.String("address", Address))
	if err := rest.echo.Start(Address); err != nil {
		rest.logger.Fatal("starting server failed", logger.Error(err))
	}
}
