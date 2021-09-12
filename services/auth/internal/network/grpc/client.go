package grpc

import (
	context "context"
	"fmt"

	"github.com/mohammadne/bookman/auth/internal/models"
	"github.com/mohammadne/bookman/auth/internal/models/pb"
	"go.opentelemetry.io/otel/trace"

	// "github.com/mohammadne/bookman/auth/internal/network/grpc/contracts"

	// "github.com/mohammadne/bookman/auth/internal/network/grpc"
	"github.com/mohammadne/bookman/auth/pkg/failures"
	"github.com/mohammadne/bookman/auth/pkg/logger"
	grpcPkg "google.golang.org/grpc"
)

type UserClient interface {
	CreateUser(context.Context, *models.UserCredential) (uint64, failures.Failure)
	GetUser(context.Context, *models.UserCredential) (uint64, failures.Failure)
}

type userClient struct {
	logger logger.Logger
	tracer trace.Tracer
	api    pb.UserClient
}

func NewUserClient(cfg *Config, lg logger.Logger, tracer trace.Tracer) (*userClient, failures.Failure) {
	client := &userClient{logger: lg, tracer: tracer}

	Address := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	connection, err := grpcPkg.Dial(Address, grpcPkg.WithInsecure())
	if err != nil {
		return nil, err
	}
	client.api = pb.NewUserClient(connection)

	return client, nil
}

func (client *userClient) CreateUser(ctx context.Context, user *models.UserCredential) (uint64, failures.Failure) {
	ctr := &pb.UserCredentialContract{Email: user.Email, Password: user.Password}
	response, err := client.api.CreateUser(ctx, ctr)

	if err != nil {
		client.logger.Error("error grpc create user", logger.Error(err))
		return 0, err
	}

	return uint64(response.Id), nil
}

func (client *userClient) GetUser(ctx context.Context, user *models.UserCredential) (uint64, failures.Failure) {
	ctr := &pb.UserCredentialContract{Email: user.Email, Password: user.Password}
	response, err := client.api.GetUser(ctx, ctr)

	if err != nil {
		client.logger.Error("error grpc get user", logger.Error(err))
		return 0, err
	}

	return uint64(response.Id), nil
}
