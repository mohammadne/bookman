package grpc

import (
	"context"
	"net"

	"github.com/mohammadne/bookman/user/internal/database"
	"github.com/mohammadne/bookman/user/internal/network/grpc/contracts"
	"github.com/mohammadne/go-pkgs/logger"
	grpc "google.golang.org/grpc"
)

type grpcServer struct {
	// injected parameters
	config   *Config
	logger   logger.Logger
	database database.Database

	// internal dependencies
	server *grpc.Server
}

func New(cfg *Config, log logger.Logger, db database.Database) *grpcServer {
	s := &grpcServer{config: cfg, logger: log, database: db}

	s.server = grpc.NewServer()
	contracts.RegisterUserServer(s.server, s)

	return s
}

func (s *grpcServer) Serve(<-chan struct{}) {
	listener, err := net.Listen("tcp", s.config.URL)
	if err != nil {
		panic(err)
	}

	s.server.Serve(listener)
}

// TODO : FIX HARD CODE
func (s *grpcServer) CreateUser(context.Context, *contracts.UserCredentialContract,
) (*contracts.UserResponse, error) {
	return &contracts.UserResponse{Id: 1}, nil
}

// TODO : FIX HARD CODE
func (s *grpcServer) GetUser(context.Context, *contracts.UserCredentialContract,
) (*contracts.UserResponse, error) {
	return &contracts.UserResponse{Id: 1}, nil
}
