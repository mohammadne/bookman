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
	config *Config
	logger logger.Logger
	tracer trace.Tracer
	client pb.AuthClient
}

func NewAuthClient(config *Config, logger logger.Logger, tracer trace.Tracer) *authClient {
	return &authClient{config: config, logger: logger, tracer: tracer}
}

func (client *authClient) Setup() {
	Address := fmt.Sprintf("%s:%s", client.config.Host, client.config.Port)
	authConnection, err := grpcPkg.Dial(Address, grpcPkg.WithInsecure())
	if err != nil {
		client.logger.Fatal(
			"error getting auth grpc connection",
			logger.String("address", Address),
			logger.Error(err),
		)
	}

	client.client = pb.NewAuthClient(authConnection)
}

var (
	errGettingToken = failures.Network{}.NewInternalServer("error getting token metadata")
	invalidToken    = failures.Network{}.NewBadRequest("invalid token")
)

func (client *authClient) GetTokenMetadata(ctx context.Context, token string) (uint64, failures.Failure) {
	ctx, span := client.tracer.Start(ctx, "database.author.get")
	defer span.End()

	contract := &pb.TokenContract{Token: token}
	response, err := client.client.TokenMetadata(ctx, contract)
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
