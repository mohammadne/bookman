package grpc

import (
	context "context"
	"net"

	"github.com/mohammadne/bookman/auth/internal/cache"
	"github.com/mohammadne/bookman/auth/internal/jwt"
	"github.com/mohammadne/bookman/auth/internal/network"
	"github.com/mohammadne/bookman/auth/internal/network/grpc/contracts"
	"github.com/mohammadne/go-pkgs/logger"
	grpc "google.golang.org/grpc"
)

type grpcServer struct {
	// injected parameters
	config *Config
	logger logger.Logger
	cache  cache.Cache
	jwt    jwt.Jwt

	// internal dependencies
	server *grpc.Server
}

func NewServer(cfg *Config, log logger.Logger, c cache.Cache, j jwt.Jwt) network.Server {
	s := &grpcServer{config: cfg, logger: log, cache: c, jwt: j}

	s.server = grpc.NewServer()
	contracts.RegisterAuthServer(s.server, s)

	return s
}

func (s *grpcServer) Serve(<-chan struct{}) {
	listener, err := net.Listen("tcp", s.config.Url)
	if err != nil {
		panic(err)
	}

	s.server.Serve(listener)
}

func (s *grpcServer) TokenMetadata(ctx context.Context, token *contracts.TokenContract,
) (*contracts.TokenMetadataResponse, error) {
	accessDetails, err := s.jwt.ExtractTokenMetadata(token.Token, jwt.Access)
	if err != nil {
		return nil, err
	}

	_, err = s.cache.GetUserId(accessDetails)
	if err != nil {
		return nil, err
	}

	return &contracts.TokenMetadataResponse{IsValid: true}, nil
}
