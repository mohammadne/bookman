package database_impl

import (
	"context"
	"fmt"

	"github.com/mohammadne/bookman/library/internal/database/ent"
	"github.com/mohammadne/bookman/library/internal/models"
)

type author interface {
	GetAuthor(ctx context.Context, id uint64) (*models.Author, error)
}

func entAuthorToModelsAuthor(entCall *ent.Author) *models.Author {
	return nil
}

func (db *database) GetAuthor(ctx context.Context, id uint64) (*models.Author, error) {
	ctx, span := db.tracer.Start(ctx, "database.author.get")
	defer span.End()

	entAuthor, err := db.client.Author.Get(ctx, int(id))
	if err != nil {
		if ent.IsNotFound(err) {
			span.RecordError(ErrNotFound)
			return nil, ErrNotFound
		}

		err = fmt.Errorf("error while getting Author from database, error: %w", err)
		span.RecordError(err)
		return nil, err
	}

	return entAuthorToModelsAuthor(entAuthor), nil
}
