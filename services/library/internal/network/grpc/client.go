package grpc

import (
	"context"
	"fmt"

	"github.com/mohammadne/bookman/library/internal/models/pb"
	"github.com/mohammadne/bookman/library/pkg/failures"
	"github.com/mohammadne/bookman/library/pkg/logger"
	"go.opentelemetry.io/otel/trace"
	grpcPkg "google.golang.org/grpc"
)

type AuthClient interface {
	GetTokenMetadata(context.Context, string) (uint64, failures.Failure)
}

type authClient struct {
	logger logger.Logger
	tracer trace.Tracer
	api    pb.AuthClient
}

func NewAuthClient(cfg *Config, lg logger.Logger, tracer trace.Tracer) (*authClient, error) {
	client := &authClient{logger: lg, tracer: tracer}

	Address := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	authConnection, err := grpcPkg.Dial(Address, grpcPkg.WithInsecure())
	if err != nil {
		return nil, err
	}
	client.api = pb.NewAuthClient(authConnection)

	return client, nil
}

var (
	errGettingToken = failures.Network{}.NewInternalServer("error getting token metadata")
	invalidToken    = failures.Network{}.NewBadRequest("invalid token")
)

func (client *authClient) GetTokenMetadata(ctx context.Context, token string) (uint64, failures.Failure) {
	ctx, span := client.tracer.Start(ctx, "database.author.get")
	defer span.End()

	contract := &pb.TokenContract{Token: token}
	response, err := client.api.TokenMetadata(ctx, contract)
	if err != nil {
		client.logger.Error("grpc get token metadata", logger.Error(err))
		span.RecordError(err)
		return 0, errGettingToken
	}

	if !response.IsValid {
		err = fmt.Errorf("not valid token metadata")
		client.logger.Error("not valid token metadata", logger.Error(err))
		span.RecordError(err)
		return 0, invalidToken
	}

	return response.Id, nil
}
