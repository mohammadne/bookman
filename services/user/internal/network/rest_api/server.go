package rest_api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/user/internal/database"
	grpc_client "github.com/mohammadne/bookman/user/internal/network/grpc/clients"
	"github.com/mohammadne/bookman/user/pkg/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type restEcho struct {
	// injected parameters
	config   *Config
	logger   logger.Logger
	database database.Database
	authGrpc grpc_client.Auth

	// internal dependencies
	echo *echo.Echo
}

func New(cfg *Config, log logger.Logger, db database.Database, ag grpc_client.Auth) *restEcho {
	handler := &restEcho{config: cfg, logger: log, database: db, authGrpc: ag}

	handler.echo = echo.New()
	handler.echo.HideBanner = true
	handler.setupRoutes()

	return handler
}

func (rest *restEcho) setupRoutes() {
	rest.echo.POST("/metrics", echo.WrapHandler(promhttp.Handler()))

	authGroup := rest.echo.Group("/users")
	authGroup.GET("/:id", rest.getUser, rest.authenticate)
	authGroup.GET("/me", rest.getMyUser, rest.authenticate)
	authGroup.GET("/search", rest.searchUsers, rest.authenticate)
}

func (rest *restEcho) Serve(<-chan struct{}) {
	Address := fmt.Sprintf("%s:%s", rest.config.Host, rest.config.Port)
	rest.logger.Info("starting server", logger.String("address", Address))
	if err := rest.echo.Start(Address); err != nil {
		rest.logger.Fatal("starting server failed", logger.Error(err))
	}
}
