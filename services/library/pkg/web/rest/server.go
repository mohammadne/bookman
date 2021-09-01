package rest

import (
	"github.com/labstack/echo/v4"
)

type server struct {
	config   *Config
	instance *echo.Echo
}

func New(cfg *Config) *server {
	server := &server{instance: echo.New(), config: cfg}
	server.instance.HideBanner = true
	return server
}

func (s *server) Serve(<-chan struct{}) error {
	return s.instance.Start(s.config.URL)
}

func (s *server) Instance() *echo.Echo {
	return s.instance
}
