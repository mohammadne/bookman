package database_impl

import (
	"github.com/mohammadne/bookman/library/internal/database/ent"
	"github.com/mohammadne/bookman/library/pkg/failures"
	"go.opentelemetry.io/otel/trace"
)

var notFoundFailure = failures.Database{}.NewNotFound("item not found")

type database struct {
	tracer trace.Tracer
	client *ent.Client
}

func New(client *ent.Client) *database {
	return &database{}
}
