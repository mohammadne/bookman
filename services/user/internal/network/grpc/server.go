package grpc

import (
	"net"

	"github.com/mohammadne/bookman/user/internal/database"
	"github.com/mohammadne/go-pkgs/logger"
	grpc "google.golang.org/grpc"
)

type Server interface {
	Serve()
}

type server struct {
	// injected parameters
	config   *Config
	logger   logger.Logger
	database database.Database

	// internal dependencies
	UnimplementedUserServer
	grpcServer *grpc.Server
}

func NewServer(cfg *Config, log logger.Logger, db database.Database) Server {
	server := &server{config: cfg, logger: log, database: db}

	server.grpcServer = grpc.NewServer()
	RegisterUserServer(server.grpcServer, server)

	return server
}

func (s *server) Serve() {
	listener, err := net.Listen("tcp", s.config.URL)
	if err != nil {
		panic(err)
	}

	s.grpcServer.Serve(listener)
}
