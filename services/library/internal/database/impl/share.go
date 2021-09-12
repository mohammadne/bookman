package database_impl

import (
	"github.com/mohammadne/bookman/library/internal/database/ent"
	"github.com/mohammadne/bookman/library/pkg/failures"
	"github.com/mohammadne/bookman/library/pkg/logger"
	"go.opentelemetry.io/otel/trace"
)

var (
	notFoundFailure = failures.Database{}.NewNotFound("item not found")
	internalFailure = failures.Database{}.NewInternalServer("error while getting item from database")
)

type database struct {
	logger logger.Logger
	tracer trace.Tracer
	client *ent.Client
}

func New(lg logger.Logger, tr trace.Tracer, client *ent.Client) *database {
	return &database{logger: lg, tracer: tr, client: client}
}
