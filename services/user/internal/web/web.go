package web

import (
	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/core/logger"
)

type WebHandler interface {
	Initialize()
	Start()
}

type echoWebHandler struct {
	// injected parameters
	config *Config
	logger *logger.Logger

	// internal dependencies
	instance *echo.Echo
}

func NewEcho(cfg *Config, log *logger.Logger) WebHandler {
	return &echoWebHandler{config: cfg, logger: log}
}

func (wh *echoWebHandler) Initialize() {
	wh.instance = echo.New()

	wh.instance.GET("/users/:id", wh.get)
}

func (wh *echoWebHandler) Start() {
	wh.instance.Start(wh.config.URL)

	(*wh.logger).Info(
		"server started",
		logger.String("address", wh.config.URL),
	)
}
