package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/user/internal/database"
	"github.com/mohammadne/go-pkgs/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server interface {
	Serve()
}

type restEcho struct {
	// injected parameters
	config   *Config
	logger   logger.Logger
	database database.Database

	// internal dependencies
	instance *echo.Echo
}

func NewEcho(cfg *Config, log logger.Logger, db database.Database) Server {
	handler := &restEcho{config: cfg, logger: log, database: db}

	handler.instance = echo.New()
	handler.instance.HideBanner = true
	handler.setupRoutes()

	return handler
}

func (rest *restEcho) setupRoutes() {
	authGroup := rest.instance.Group("/users")

	authGroup.POST("/metrics", echo.WrapHandler(promhttp.Handler()))
	authGroup.POST("/:id", rest.getUser)
	authGroup.POST("/me", rest.getMe)
}

func (rest *restEcho) Serve() {
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
