package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/mohammadne/bookman/user/internal/models"
	"github.com/mohammadne/bookman/user/internal/models/pb"
	"github.com/mohammadne/bookman/user/internal/storage"
	"github.com/mohammadne/bookman/user/pkg/failures"
	"github.com/mohammadne/bookman/user/pkg/logger"

	"google.golang.org/grpc"
)

type grpcServer struct {
	// injected parameters
	config  *Config
	logger  logger.Logger
	storage storage.Storage

	// internal dependencies
	server *grpc.Server
	pb.UnimplementedUserServer
}

func NewServer(cfg *Config, log logger.Logger, storage storage.Storage) *grpcServer {
	s := &grpcServer{config: cfg, logger: log, storage: storage}

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

func (server *grpcServer) CreateUser(_ context.Context, credentials *pb.UserCredentialContract,
) (*pb.UserResponse, error) {
	user, failure := server.storage.FindUserByEmail(credentials.Email)
	if user != nil || failure == nil {
		return nil, failures.Network{}.NewBadRequest("email is already registered")
	}

	user = &models.User{Email: credentials.Email, Password: credentials.Password}
	failure = server.storage.CreateUser(user)
	if failure != nil {
		return nil, failure
	}

	return &pb.UserResponse{Id: user.Id}, nil
}

func (server *grpcServer) GetUser(_ context.Context, credentials *pb.UserCredentialContract,
) (*pb.UserResponse, error) {
	user, failure := server.storage.FindUserByEmailAndPassword(credentials.Email, credentials.Password)
	if failure != nil {
		return nil, failure
	}

	return &pb.UserResponse{Id: user.Id}, nil
}
