package grpc_client

import (
	"context"
	"errors"

	"github.com/mohammadne/bookman/user/internal/models"
	"github.com/mohammadne/bookman/user/internal/network/grpc/contracts"
	"github.com/mohammadne/go-pkgs/logger"
)

var (
	errInvalidToken = errors.New("token is invalid")
)

type AuthClient interface {
	GetTokenMetadata(string) (*models.User, error)
}

func (g *grpcClient) GetTokenMetadata(token string) (*models.User, error) {
	contract := &contracts.TokenContract{Token: token}
	response, err := g.authClient.TokenMetadata(context.Background(), contract)
	if err != nil {
		g.logger.Error("grpc get token metadata", logger.Error(err))
		return nil, err
	}

	if response.IsValid {
		g.logger.Error("not valid token metadata", logger.Error(err))
		return nil, errInvalidToken
	}

	return &models.User{Id: response.Id}, nil
}
