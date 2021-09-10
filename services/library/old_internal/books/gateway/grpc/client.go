package grpc

import (
	"context"

	"github.com/mohammadne/bookman/library/internal/books/gateway/grpc/contracts"
	"github.com/mohammadne/bookman/library/pkg/failures"
	"github.com/mohammadne/bookman/library/pkg/logger"
	"github.com/mohammadne/bookman/library/pkg/web/grpc"
)

type AuthClient interface {
	GetTokenMetadata(string) (uint64, failures.Failure)
}

type authClient struct {
	logger   logger.Logger
	instance contracts.AuthClient
}

func NewAuthClient(cfg *grpc.Config, lg logger.Logger) *authClient {
	connection, err := grpc.NewClient(cfg)
	if err != nil {
		lg.Error("error establish connection", logger.Error(err))
		return nil
	}

	return &authClient{
		logger:   lg,
		instance: contracts.NewAuthClient(connection),
	}
}

var (
	errGettingToken = failures.Web{}.NewInternalServer("error getting token metadata")
	invalidToken    = failures.Web{}.NewInternalServer("invalid token")
)

func (client *authClient) GetTokenMetadata(token string) (uint64, failures.Failure) {
	contract := &contracts.TokenContract{Token: token}
	response, err := client.instance.TokenMetadata(context.Background(), contract)
	if err != nil {
		client.logger.Error("grpc get token metadata", logger.Error(err))
		return 0, errGettingToken
	}

	if !response.IsValid {
		client.logger.Error("not valid token metadata", logger.Error(err))
		return 0, invalidToken
	}

	return response.Id, nil
}
