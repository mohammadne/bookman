package database_impl

import (
	"context"
	"io"

	"github.com/mohammadne/bookman/library/internal/database/ent/migrate"
)

type Migration interface {
	MigratePreview(ctx context.Context, writer io.Writer) error
	Migrate(ctx context.Context) error
}

func (db *database) MigratePreview(ctx context.Context, writer io.Writer) error {
	return db.client.Schema.WriteTo(ctx,
		writer,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
}

func (db *database) Migrate(ctx context.Context) error {
	return db.client.Schema.Create(ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
}
