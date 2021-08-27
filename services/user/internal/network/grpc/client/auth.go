package grpc_client

import (
	"context"

	"github.com/mohammadne/bookman/user/internal/network/grpc/contracts"
	"github.com/mohammadne/go-pkgs/logger"
)

type AuthClient interface {
	GetTokenMetadata(string) (*contracts.TokenMetadataResponse, error)
}

func (g *grpcClient) GetTokenMetadata(token string) (*contracts.TokenMetadataResponse, error) {
	contract := &contracts.TokenContract{Token: token}
	response, err := g.authClient.TokenMetadata(context.Background(), contract)
	if err != nil {
		g.logger.Error("error grpc get token metadata", logger.Error(err))
		return nil, err
	}

	return response, nil
}
