package grpc_client

import (
	"github.com/mohammadne/bookman/user/internal/network/grpc/contracts"
	"github.com/mohammadne/go-pkgs/logger"
	"google.golang.org/grpc"
)

type Client interface {
	AuthClient
}

type grpcClient struct {
	config *Config
	logger logger.Logger

	authClient contracts.AuthClient
}

func New(cfg *Config, log logger.Logger) Client {
	return &grpcClient{config: cfg, logger: log}
}

func (g *grpcClient) Setup() {
	authConnection, err := grpc.Dial(g.config.AuthAddress, grpc.WithInsecure())
	if err != nil {
		g.logger.Fatal(
			"error getting auth grpc connection",
			logger.String("address", g.config.AuthAddress),
			logger.Error(err),
		)
	}

	g.authClient = contracts.NewAuthClient(authConnection)
}
