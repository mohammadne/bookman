package grpc

import (
	context "context"
	"fmt"

	"github.com/mohammadne/bookman/auth/internal/models"
	"github.com/mohammadne/bookman/auth/internal/models/pb"
	"go.opentelemetry.io/otel/trace"

	// "github.com/mohammadne/bookman/auth/internal/network/grpc/contracts"

	// "github.com/mohammadne/bookman/auth/internal/network/grpc"
	"github.com/mohammadne/bookman/auth/pkg/logger"
	grpcPkg "google.golang.org/grpc"
)

type UserClient interface {
	CreateUser(*models.UserCredential) (uint64, error)
	GetUser(*models.UserCredential) (uint64, error)
}

type userClient struct {
	logger logger.Logger
	tracer trace.Tracer
	api    pb.UserClient
}

func NewUserClient(cfg *Config, lg logger.Logger, tracer trace.Tracer) (*userClient, error) {
	client := &userClient{logger: lg, tracer: tracer}

	Address := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	connection, err := grpcPkg.Dial(Address, grpcPkg.WithInsecure())
	if err != nil {
		return nil, err
	}
	client.api = pb.NewUserClient(connection)

	return client, nil
}

func (client *userClient) CreateUser(user *models.UserCredential) (uint64, error) {
	response, err := client.api.CreateUser(
		context.Background(),
		&pb.UserCredentialContract{Email: user.Email, Password: user.Password},
	)

	if err != nil {
		client.logger.Error("error grpc create user", logger.Error(err))
		return 0, err
	}

	return uint64(response.Id), nil
}

func (client *userClient) GetUser(user *models.UserCredential) (uint64, error) {
	response, err := client.api.GetUser(
		context.Background(),
		&pb.UserCredentialContract{Email: user.Email, Password: user.Password},
	)

	if err != nil {
		client.logger.Error("error grpc get user", logger.Error(err))
		return 0, err
	}

	return uint64(response.Id), nil
}
