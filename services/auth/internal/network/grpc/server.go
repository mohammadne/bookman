package grpc

import (
	context "context"
	"net"

	"github.com/mohammadne/bookman/auth/internal/cache"
	"github.com/mohammadne/bookman/auth/internal/jwt"
	"github.com/mohammadne/go-pkgs/logger"
	grpc "google.golang.org/grpc"
)

type Server interface {
	Serve()
}

type server struct {
	// injected parameters
	config *Config
	logger logger.Logger
	cache  cache.Cache
	jwt    jwt.Jwt

	// internal dependencies
	UnimplementedAuthServer
	grpcServer *grpc.Server
}

func NewServer(cfg *Config, log logger.Logger, c cache.Cache, j jwt.Jwt) Server {
	server := &server{config: cfg, logger: log, cache: c, jwt: j}

	server.grpcServer = grpc.NewServer()
	RegisterAuthServer(server.grpcServer, server)

	return server
}

func (s *server) Serve() {
	listener, err := net.Listen("tcp", s.config.URL)
	if err != nil {
		panic(err)
	}

	s.grpcServer.Serve(listener)
}

func (s *server) TokenMetadata(context.Context, *TokenContract) (*TokenMetadataResponse, error) {
	return nil, nil
}
