package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/core/logger"
	"github.com/mohammadne/bookman/user/internal/database"
)

type RestAPI interface {
	Start()
	StartG()
}

type echoRestAPI struct {
	// injected parameters
	config   *Config
	logger   logger.Logger
	database database.Database

	// internal dependencies
	instance *echo.Echo
}

func NewEcho(cfg *Config, log logger.Logger, db database.Database) RestAPI {
	handler := &echoRestAPI{config: cfg, logger: log, database: db}

	handler.instance = echo.New()
	handler.instance.HideBanner = true

	handler.instance.GET("/users/:id", handler.get)
	handler.instance.GET("/users/me", handler.getMe)

	return handler
}

func (rest *echoRestAPI) Start() {
	rest.logger.Info(
		"starting server",
		logger.String("address", rest.config.URL),
	)

	if err := rest.instance.Start(rest.config.URL); err != nil {
		rest.logger.Fatal("starting server failed", logger.Error(err))
	}
}

func (rest *echoRestAPI) StartG() {
	go func() { rest.Start() }()
}
