package grpc

import (
	context "context"

	"github.com/mohammadne/bookman/user/internal/network/grpc/contracts"

	"github.com/mohammadne/bookman/user/internal/network/grpc"
	"github.com/mohammadne/go-pkgs/failures"
	"github.com/mohammadne/go-pkgs/logger"
	grpcPkg "google.golang.org/grpc"
)

type Auth interface {
	GetTokenMetadata(string) (uint64, failures.Failure)
}

type authClient struct {
	logger logger.Logger
	config *grpc.Config

	client contracts.AuthClient
}

func NewUser(cfg *grpc.Config, log logger.Logger) *authClient {
	return &authClient{config: cfg, logger: log}
}

func (g *authClient) Setup() {
	authConnection, err := grpcPkg.Dial(g.config.Url, grpcPkg.WithInsecure())
	if err != nil {
		g.logger.Fatal(
			"error getting auth grpc connection",
			logger.String("address", g.config.Url),
			logger.Error(err),
		)
	}

	g.client = contracts.NewAuthClient(authConnection)
}

var (
	errGettingToken = failures.Rest{}.NewInternalServer("error getting token metadata")
	invalidToken    = failures.Rest{}.NewInternalServer("invalid token")
)

func (g *authClient) GetTokenMetadata(token string) (uint64, failures.Failure) {
	contract := &contracts.TokenContract{Token: token}
	response, err := g.client.TokenMetadata(context.Background(), contract)
	if err != nil {
		g.logger.Error("grpc get token metadata", logger.Error(err))
		return 0, errGettingToken
	}

	if !response.IsValid {
		g.logger.Error("not valid token metadata", logger.Error(err))
		return 0, invalidToken
	}

	return response.Id, nil
}
