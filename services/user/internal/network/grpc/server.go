package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/mohammadne/bookman/user/internal/database"
	"github.com/mohammadne/bookman/user/internal/models"
	"github.com/mohammadne/bookman/user/internal/models/pb"
	"github.com/mohammadne/bookman/user/pkg/logger"
	"github.com/mohammadne/go-pkgs/failures"
	"google.golang.org/grpc"
)

type grpcServer struct {
	// injected parameters
	config   *Config
	logger   logger.Logger
	database database.Database

	// internal dependencies
	server *grpc.Server
	pb.UnimplementedUserServer
}

func NewServer(cfg *Config, log logger.Logger, db database.Database) *grpcServer {
	s := &grpcServer{config: cfg, logger: log, database: db}

	s.server = grpc.NewServer()
	pb.RegisterUserServer(s.server, s)

	return s
}

func (s *grpcServer) Serve(<-chan struct{}) {
	address := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	s.server.Serve(listener)
}

func (s *grpcServer) CreateUser(_ context.Context, credentials *pb.UserCredentialContract,
) (*pb.UserResponse, error) {
	user, failure := s.database.FindUserByEmail(credentials.Email)
	if user != nil || failure == nil {
		return nil, failures.Rest{}.NewBadRequest("email is already registered")
	}

	user = &models.User{Email: credentials.Email, Password: credentials.Password}
	failure = s.database.CreateUser(user)
	if failure != nil {
		return nil, failure
	}

	return &pb.UserResponse{Id: user.Id}, nil
}

func (s *grpcServer) GetUser(_ context.Context, credentials *pb.UserCredentialContract,
) (*pb.UserResponse, error) {
	user, failure := s.database.FindUserByEmailAndPassword(credentials.Email, credentials.Password)
	if failure != nil {
		return nil, failure
	}

	return &pb.UserResponse{Id: user.Id}, nil
}
