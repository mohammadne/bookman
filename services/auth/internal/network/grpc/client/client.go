package grpc_client

import (
	"github.com/mohammadne/bookman/auth/internal/network/grpc/contracts"
	"github.com/mohammadne/go-pkgs/logger"
	"google.golang.org/grpc"
)

type Client interface {
	UserClient
}

type grpcClient struct {
	config *Config
	logger logger.Logger

	userClient contracts.UserClient
}

func New(cfg *Config, log logger.Logger) Client {
	return &grpcClient{config: cfg, logger: log}
}

func (g *grpcClient) Setup() {
	userConnection, err := grpc.Dial(g.config.UserAddress, grpc.WithInsecure())
	if err != nil {
		g.logger.Fatal(
			"error getting user grpc connection",
			logger.String("address", g.config.UserAddress),
			logger.Error(err),
		)
	}

	g.userClient = contracts.NewUserClient(userConnection)
}
