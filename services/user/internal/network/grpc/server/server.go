package grpc

import (
	"context"
	"net"

	"github.com/mohammadne/bookman/user/internal/database"
	"github.com/mohammadne/bookman/user/internal/models"
	"github.com/mohammadne/bookman/user/internal/network/grpc/contracts"
	"github.com/mohammadne/go-pkgs/failures"
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
	contracts.UnimplementedUserServer
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

func (s *grpcServer) CreateUser(_ context.Context, credentials *contracts.UserCredentialContract,
) (*contracts.UserResponse, error) {
	user, failure := s.database.FindUserByEmail(credentials.Email)
	if user != nil || failure == nil {
		return nil, failures.Rest{}.NewBadRequest("email is already registered")
	}

	user = &models.User{}
	failure = s.database.CreateUser(user)
	if failure != nil {
		return nil, failure
	}

	return &contracts.UserResponse{Id: user.Id}, nil
}

func (s *grpcServer) GetUser(_ context.Context, credentials *contracts.UserCredentialContract,
) (*contracts.UserResponse, error) {
	user, failure := s.database.FindUserByEmailAndPassword(credentials.Email, credentials.Password)
	if failure != nil {
		return nil, failure
	}

	return &contracts.UserResponse{Id: user.Id}, nil
}
