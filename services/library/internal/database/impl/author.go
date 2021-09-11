package database_impl

import (
	"context"

	"github.com/mohammadne/bookman/library/internal/database/ent"
	"github.com/mohammadne/bookman/library/internal/models"
	"github.com/mohammadne/bookman/library/pkg/failures"
)

type Author interface {
	GetAuthor(ctx context.Context, id uint64) (*models.Author, failures.Failure)
}

func entAuthorToModelsAuthor(entCall *ent.Author) *models.Author {
	return nil
}

func (db *database) GetAuthor(ctx context.Context, id uint64) (*models.Author, failures.Failure) {
	ctx, span := db.tracer.Start(ctx, "database.author.get")
	defer span.End()

	entAuthor, err := db.client.Author.Get(ctx, int(id))
	if err != nil {
		if ent.IsNotFound(err) {
			span.RecordError(notFoundFailure)
			return nil, notFoundFailure
		}

		errStr := "error while getting Author from database"
		failure := failures.Database{}.NewInternalServer(errStr)
		span.RecordError(err)
		return nil, failure
	}

	return entAuthorToModelsAuthor(entAuthor), nil
}
