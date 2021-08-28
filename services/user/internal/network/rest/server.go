package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/user/internal/database"
	"github.com/mohammadne/go-pkgs/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type restEcho struct {
	// injected parameters
	config   *Config
	logger   logger.Logger
	database database.Database

	// internal dependencies
	instance *echo.Echo
}

func New(cfg *Config, log logger.Logger, db database.Database) *restEcho {
	handler := &restEcho{config: cfg, logger: log, database: db}

	handler.instance = echo.New()
	handler.instance.HideBanner = true
	handler.setupRoutes()

	return handler
}

func (rest *restEcho) setupRoutes() {
	authGroup := rest.instance.Group("/users")

	authGroup.POST("/metrics", echo.WrapHandler(promhttp.Handler()))
	authGroup.GET("/:id", rest.getUser)
	authGroup.GET("/me", rest.getMe)
}

func (rest *restEcho) Serve(<-chan struct{}) {
	rest.logger.Info(
		"starting server",
		logger.String("address", rest.config.URL),
	)

	if err := rest.instance.Start(rest.config.URL); err != nil {
		rest.logger.Fatal("starting server failed", logger.Error(err))
	}
}
