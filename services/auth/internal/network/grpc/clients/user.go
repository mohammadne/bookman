package grpc_clients

import (
	context "context"

	"github.com/mohammadne/bookman/auth/internal/models"
	"github.com/mohammadne/bookman/auth/internal/network/grpc/contracts"

	"github.com/mohammadne/bookman/auth/internal/network/grpc"
	"github.com/mohammadne/go-pkgs/logger"
	grpcPkg "google.golang.org/grpc"
)

type User interface {
	CreateUser(*models.UserCredential) (uint64, error)
	GetUser(*models.UserCredential) (uint64, error)
}

type userClient struct {
	logger logger.Logger
	config *grpc.Config

	client contracts.UserClient
}

func NewUser(cfg *grpc.Config, log logger.Logger) *userClient {
	// grpc.Config{}
	return &userClient{config: cfg, logger: log}
}

func (g *userClient) Setup() {
	userConnection, err := grpcPkg.Dial(g.config.Url, grpcPkg.WithInsecure())
	if err != nil {
		g.logger.Fatal(
			"error getting user grpc connection",
			logger.String("address", g.config.Url),
			logger.Error(err),
		)
	}

	g.client = contracts.NewUserClient(userConnection)
}

func (c *userClient) CreateUser(user *models.UserCredential) (uint64, error) {
	response, err := c.client.CreateUser(
		context.Background(),
		&contracts.UserCredentialContract{Email: user.Email, Password: user.Password},
	)

	if err != nil {
		c.logger.Error("error grpc create user", logger.Error(err))
		return 0, err
	}

	return uint64(response.Id), nil
}

func (c *userClient) GetUser(user *models.UserCredential) (uint64, error) {
	response, err := c.client.GetUser(
		context.Background(),
		&contracts.UserCredentialContract{Email: user.Email, Password: user.Password},
	)

	if err != nil {
		c.logger.Error("error grpc get user", logger.Error(err))
		return 0, err
	}

	return uint64(response.Id), nil
}
