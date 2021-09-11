package rest_api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/library/internal/database"
	"github.com/mohammadne/bookman/library/internal/network/grpc"
	"github.com/mohammadne/bookman/library/pkg/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/trace"
)

type restApi struct {
	config   *Config
	logger   logger.Logger
	tracer   trace.Tracer
	database database.Database
	authGrpc grpc.AuthClient

	// internal dependencies
	echo *echo.Echo
}

func NewServer(cfg *Config, lg logger.Logger, tr trace.Tracer, db database.Database, ac grpc.AuthClient) *restApi {
	rest := &restApi{config: cfg, logger: lg, tracer: tr, database: db, authGrpc: ac}

	rest.echo = echo.New()
	rest.echo.HideBanner = true
	rest.setupRoutes()

	return rest
}

func (rest *restApi) setupRoutes() {
	rest.echo.POST("/metrics", echo.WrapHandler(promhttp.Handler()))

	booksGroup := rest.echo.Group("/books")
	booksGroup.GET("/:id", rest.getBook, rest.authenticate)

	authorsGroup := rest.echo.Group("/authors")
	authorsGroup.GET("/:id", rest.getAuthor, rest.authenticate)
}

func (rest *restApi) Serve(<-chan struct{}) {
	address := fmt.Sprintf("%s:%s", rest.config.Host, rest.config.Port)
	rest.logger.Info("starting server", logger.String("address", address))
	if err := rest.echo.Start(address); err != nil {
		rest.logger.Fatal("starting server failed", logger.Error(err))
	}
}
