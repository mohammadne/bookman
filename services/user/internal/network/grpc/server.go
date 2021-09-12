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
	"go.opentelemetry.io/otel/trace"

	"google.golang.org/grpc"
)

type grpcServer struct {
	config  *Config
	logger  logger.Logger
	tracer  trace.Tracer
	storage storage.Storage

	// internal dependencies
	server *grpc.Server
	pb.UnimplementedUserServer
}

func NewServer(cfg *Config, log logger.Logger, t trace.Tracer, storage storage.Storage) *grpcServer {
	s := &grpcServer{config: cfg, logger: log, tracer: t, storage: storage}

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

func (server *grpcServer) CreateUser(ctx context.Context, credentials *pb.UserCredentialContract,
) (*pb.UserResponse, error) {
	ctx, span := server.tracer.Start(ctx, "network.grpc.server.create_user")
	defer span.End()

	user, failure := server.storage.FindUserByEmail(ctx, credentials.Email)
	if user != nil || failure == nil {
		failure := failures.Network{}.NewBadRequest("email is already registered")
		span.RecordError(failure)
		return nil, failure
	}

	user = &models.User{Email: credentials.Email, Password: credentials.Password}
	failure = server.storage.CreateUser(ctx, user)
	if failure != nil {
		span.RecordError(failure)
		return nil, failure
	}

	return &pb.UserResponse{Id: user.Id}, nil
}

func (server *grpcServer) GetUser(ctx context.Context, credentials *pb.UserCredentialContract,
) (*pb.UserResponse, error) {
	ctx, span := server.tracer.Start(ctx, "network.grpc.server.get_user")
	defer span.End()

	user, failure := server.storage.FindUserByEmailAndPassword(ctx, credentials.Email, credentials.Password)
	if failure != nil {
		span.RecordError(failure)
		return nil, failure
	}

	return &pb.UserResponse{Id: user.Id}, nil
}
