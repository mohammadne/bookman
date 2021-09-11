package database_impl

import (
	"errors"

	"github.com/mohammadne/bookman/library/internal/database/ent"
	"go.opentelemetry.io/otel/trace"
)

var ErrNotFound = errors.New("item not found")

type database struct {
	tracer trace.Tracer
	client *ent.Client
}

func New(client *ent.Client) *database {
	return &database{}
}
