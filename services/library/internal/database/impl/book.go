package database_impl

import (
	"context"

	"github.com/mohammadne/bookman/library/internal/database/ent"
	"github.com/mohammadne/bookman/library/internal/models"
	"github.com/mohammadne/bookman/library/pkg/failures"
)

type Book interface {
	GetBook(ctx context.Context, id uint64) (*models.Book, failures.Failure)
}

func entBookToModelsBook(entCall *ent.Book) *models.Book {
	return nil
}

func (db *database) GetBook(ctx context.Context, id uint64) (*models.Book, failures.Failure) {
	ctx, span := db.tracer.Start(ctx, "database.book.get")
	defer span.End()

	entBook, err := db.client.Book.Get(ctx, int(id))
	if err != nil {
		if ent.IsNotFound(err) {
			span.RecordError(notFoundFailure)
			return nil, notFoundFailure
		}

		span.RecordError(err)
		return nil, internalFailure
	}

	return entBookToModelsBook(entBook), nil
}
