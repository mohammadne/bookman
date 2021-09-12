package grpc

import (
	context "context"
	"fmt"
	"net"

	"github.com/mohammadne/bookman/auth/internal/cache"
	"github.com/mohammadne/bookman/auth/internal/jwt"
	"github.com/mohammadne/bookman/auth/internal/models/pb"
	"github.com/mohammadne/bookman/auth/internal/network"
	"github.com/mohammadne/go-pkgs/logger"
	"google.golang.org/grpc"
)

type grpcServer struct {
	// injected parameters
	config *Config
	logger logger.Logger
	cache  cache.Cache
	jwt    jwt.Jwt

	// internal dependencies
	server *grpc.Server
	pb.UnimplementedAuthServer
}

func NewServer(cfg *Config, log logger.Logger, c cache.Cache, j jwt.Jwt) network.Server {
	s := &grpcServer{config: cfg, logger: log, cache: c, jwt: j}

	s.server = grpc.NewServer()
	pb.RegisterAuthServer(s.server, s)

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

func (s *grpcServer) TokenMetadata(ctx context.Context, token *pb.TokenContract,
) (*pb.TokenMetadataResponse, error) {
	accessDetails, err := s.jwt.ExtractTokenMetadata(token.Token, jwt.Access)
	if err != nil {
		return nil, err
	}

	userId, err := s.cache.GetUserId(accessDetails)
	if err != nil {
		return nil, err
	}

	return &pb.TokenMetadataResponse{IsValid: true, Id: userId}, nil
}
