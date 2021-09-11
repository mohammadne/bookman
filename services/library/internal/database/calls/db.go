package database_impl

import (
	"github.com/mohammadne/bookman/library/internal/database/ent"
	"go.opentelemetry.io/otel/trace"
)

type Database interface {
	migration
	book
	author
}

type database struct {
	tracer trace.Tracer
	client *ent.Client
}

func New(client *ent.Client) Database {
	return &database{}
}
