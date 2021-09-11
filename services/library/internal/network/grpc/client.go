package grpc

import (
	"github.com/mohammadne/bookman/library/internal/models/pb"
	"github.com/mohammadne/bookman/library/pkg/failures"
	"github.com/mohammadne/bookman/library/pkg/logger"
	"github.com/mohammadne/bookman/library/pkg/web/grpc"
	"go.opentelemetry.io/otel/trace"
)

type AuthClient interface {
	GetTokenMetadata(string) (uint64, failures.Failure)
}

type authClient struct {
	logger logger.Logger
	tracer trace.Tracer
	config *grpc.Config
	client pb.AuthClient
}
