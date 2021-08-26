package grpc_server

import (
	context "context"
	"net"

	"github.com/mohammadne/bookman/auth/internal/cache"
	"github.com/mohammadne/bookman/auth/internal/jwt"
	"github.com/mohammadne/bookman/auth/internal/network"
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
	UnimplementedAuthServer
	server *grpc.Server
}

func New(cfg *Config, log logger.Logger, c cache.Cache, j jwt.Jwt) network.Server {
	s := &grpcServer{config: cfg, logger: log, cache: c, jwt: j}

	s.server = grpc.NewServer()
	RegisterAuthServer(s.server, s)

	return s
}

func (s *grpcServer) Serve(<-chan struct{}) {
	listener, err := net.Listen("tcp", s.config.URL)
	if err != nil {
		panic(err)
	}

	s.server.Serve(listener)
}

func (s *grpcServer) TokenMetadata(context.Context, *TokenContract) (*TokenMetadataResponse, error) {
	return nil, nil
}
