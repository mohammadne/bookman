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

func (client *userClient) CreateUser(ctx context.Context, user *models.UserCredential) (uint64, failures.Failure) {
	ctx, span := client.tracer.Start(ctx, "network.grpc.client.user.create_user")
	defer span.End()

	ctr := &pb.UserCredentialContract{Email: user.Email, Password: user.Password}
	operation := func() (*pb.UserResponse, error) {
		return client.api.CreateUser(ctx, ctr)
	}

	id, failure := client.operate(operation)
	if failure != nil {
		client.logger.Error(failure.Message(), logger.Error(failure))
		span.RecordError(failure)
		return 0, failure
	}

	return id, nil
}

func (client *userClient) GetUser(ctx context.Context, user *models.UserCredential) (uint64, failures.Failure) {
	ctx, span := client.tracer.Start(ctx, "network.grpc.client.user.get_user")
	defer span.End()

	ctr := &pb.UserCredentialContract{Email: user.Email, Password: user.Password}
	operation := func() (*pb.UserResponse, error) {
		return client.api.GetUser(ctx, ctr)
	}

	id, failure := client.operate(operation)
	if failure != nil {
		client.logger.Error(failure.Message(), logger.Error(failure))
		span.RecordError(failure)
		return 0, failure
	}

	return id, nil
}

func (client *userClient) operate(operation func() (*pb.UserResponse, error)) (uint64, failures.Failure) {
	response, err := operation()

	if err != nil {
		return 0, failures.Network{}.NewInternalServer(err.Error())
	}

	if response.Id == 0 {
		return 0, failures.Network{}.NewBadRequest("invalid user response")
	}

	return uint64(response.Id), nil
}
