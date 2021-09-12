package database_impl

import (
	"context"
	"io"

	"github.com/mohammadne/bookman/library/internal/database/ent/migrate"
)

type Migration interface {
	MigratePreview(writer io.Writer) error
	Migrate() error
}

func (db *database) MigratePreview(writer io.Writer) error {
	return db.client.Schema.WriteTo(
		context.TODO(),
		writer,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
}

func (db *database) Migrate() error {
	return db.client.Schema.Create(
		context.TODO(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
}
