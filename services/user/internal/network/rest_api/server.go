package rest_api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/user/internal/network/grpc"
	"github.com/mohammadne/bookman/user/internal/storage"
	"github.com/mohammadne/bookman/user/pkg/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/trace"
)

type restServer struct {
	// injected parameters
	config   *Config
	logger   logger.Logger
	tracer   trace.Tracer
	storage  storage.Storage
	authGrpc grpc.AuthClient

	// internal dependencies
	echo *echo.Echo
}

func New(cfg *Config, lg logger.Logger, t trace.Tracer, s storage.Storage, ac grpc.AuthClient) *restServer {
	handler := &restServer{config: cfg, logger: lg, tracer: t, storage: s, authGrpc: ac}

	handler.echo = echo.New()
	handler.echo.HideBanner = true
	handler.setupRoutes()

	return handler
}

func (rest *restServer) setupRoutes() {
	rest.echo.POST("/metrics", echo.WrapHandler(promhttp.Handler()))

	userGroup := rest.echo.Group("/users")
	userGroup.GET("/:id", rest.getUser, rest.authenticate)
	userGroup.GET("/me", rest.getMyUser, rest.authenticate)
}

func (rest *restServer) Serve(<-chan struct{}) {
	Address := fmt.Sprintf("%s:%s", rest.config.Host, rest.config.Port)
	rest.logger.Info("starting server", logger.String("address", Address))
	if err := rest.echo.Start(Address); err != nil {
		rest.logger.Fatal("starting server failed", logger.Error(err))
	}
}
