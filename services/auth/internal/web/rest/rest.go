package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/core/logger"
)

type Rest interface {
	Start()
}

type echoRestAPI struct {
	// injected parameters
	config *Config
	logger logger.Logger

	// internal dependencies
	instance *echo.Echo
}
