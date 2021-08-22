package web

import (
	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/core/logger"
	"github.com/mohammadne/bookman/user/internal/database"
)

type Web interface {
	Start()
	StartG()
}

type echoWebHandler struct {
	// injected parameters
	config   *Config
	logger   logger.Logger
	database database.Database

	// internal dependencies
	instance *echo.Echo
}

func NewEcho(cfg *Config, log logger.Logger, db database.Database) Web {
	handler := &echoWebHandler{config: cfg, logger: log, database: db}

	handler.instance = echo.New()
	handler.instance.HideBanner = true

	handler.instance.GET("/users/:id", handler.get)
	handler.instance.GET("/users/me", handler.getMe)

	return handler
}

func (wh *echoWebHandler) Start() {
	wh.logger.Info(
		"starting server",
		logger.String("address", wh.config.URL),
	)

	if err := wh.instance.Start(wh.config.URL); err != nil {
		wh.logger.Fatal("starting server failed", logger.Error(err))
	}
}

func (wh *echoWebHandler) StartG() {
	go func() { wh.Start() }()
}
