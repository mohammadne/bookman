package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mohammadne/bookman/library/internal/database/ent"
	impl "github.com/mohammadne/bookman/library/internal/database/impl"
	"github.com/mohammadne/bookman/library/pkg/logger"

	"go.opentelemetry.io/otel/trace"
)

type Database interface {
	impl.Author
	impl.Book
	impl.Migration
}

// NewClient creates an ent database connection based on entry config.
func NewClient(config *Config, lg logger.Logger, tr trace.Tracer) (Database, error) {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		config.User, config.Password, config.Host, config.DatabaseName,
	)

	client, err := ent.Open(config.Driver, dataSourceName)
	return impl.New(lg, tr, client), err
}
