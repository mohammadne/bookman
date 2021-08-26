package grpc_client

import (
	"context"

	"github.com/mohammadne/bookman/auth/internal/models"
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
			"error getting grpc connection",
			logger.String("address", g.config.UserAddress),
			logger.Error(err),
		)
	}

	g.userClient = contracts.NewUserClient(userConnection)
}

type UserClient interface {
	CreateUser(models.UserCredential) (*contracts.UserResponse, error)
	GetUser(models.UserCredential) (*contracts.UserResponse, error)
}

func (g *grpcClient) CreateUser(user models.UserCredential) (*contracts.UserResponse, error) {
	// in *user.UserCredentialContract
	response, err := g.userClient.CreateUser(context.Background(), nil)
	if err != nil {
		g.logger.Error("error grpc create user", logger.Error(err))
		return nil, err
	}

	return response, nil
}

func (g *grpcClient) GetUser(user models.UserCredential) (*contracts.UserResponse, error) {
	// in *user.UserCredentialContract
	response, err := g.userClient.GetUser(context.Background(), nil)
	if err != nil {
		g.logger.Error("error grpc get user", logger.Error(err))
		return nil, err
	}

	return response, nil
}
