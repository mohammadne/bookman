package rest

import (
	"github.com/labstack/echo/v4"
)

type Server interface {
	Serve(<-chan struct{})
	Instance() *echo.Echo
}

type server struct {
	config   *Config
	instance *echo.Echo
}

func New(cfg *Config) Server {
	return &server{instance: echo.New(), config: cfg}
}

func (s *server) Serve(<-chan struct{}) {
	if err := s.instance.Start(s.config.URL); err != nil {
		panic("server failed to start")
	}
}

func (s *server) Instance() *echo.Echo {
	return s.instance
}
