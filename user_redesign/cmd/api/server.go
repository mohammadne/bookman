package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/user/config"
	delivery "github.com/mohammadne/bookman/user/internal/delivery/http"
	"github.com/mohammadne/bookman/user/internal/repositories"
	"github.com/mohammadne/bookman/user/internal/usecases"
	"github.com/mohammadne/bookman/user/pkg/logger"
)

const (
	certFile       = "ssl/Server.crt"
	keyFile        = "ssl/Server.pem"
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
)

// Server struct
type Server struct {
	echo   *echo.Echo
	cfg    *config.Config
	db     *sqlx.DB
	logger logger.Logger
}

// NewServer New Server constructor
func NewServer(cfg *config.Config, db *sqlx.DB, redisClient, logger logger.Logger) *Server {
	return &Server{echo: echo.New(), cfg: cfg, db: db, logger: logger}
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:           s.cfg.Server.Port,
		ReadTimeout:    time.Second * s.cfg.Server.ReadTimeout,
		WriteTimeout:   time.Second * s.cfg.Server.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		s.logger.Infof("Server is listening on PORT: %s", s.cfg.Server.Port)
		if err := s.echo.StartServer(server); err != nil {
			s.logger.Fatalf("Error starting Server: ", err)
		}
	}()

	go func() {
		s.logger.Infof("Starting Debug Server on PORT: %s", s.cfg.Server.PprofPort)
		if err := http.ListenAndServe(s.cfg.Server.PprofPort, http.DefaultServeMux); err != nil {
			s.logger.Errorf("Error PPROF ListenAndServe: %s", err)
		}
	}()

	if err := s.MapHandlers(s.echo); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	s.logger.Info("Server Exited Properly")
	return s.echo.Server.Shutdown(ctx)
}

func (s *Server) MapHandlers(e *echo.Echo) error {
	// Init repositories
	userR := repositories.NewUserRepository(s.db)

	// Init usecases
	userU := usecases.NewUserUsecase(userR)

	// Init handlers
	userH := delivery.NewUserHandler(userU)

	// WEB framework path prefix
	v1 := e.Group("/api/v1")
	userGroup := v1.Group("/users")

	// handle routes
	delivery.MapUserRoutes(userGroup, userH)

	return nil
}
