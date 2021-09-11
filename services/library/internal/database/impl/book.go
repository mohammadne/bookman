package database_impl

import (
	"context"
	"fmt"

	"github.com/mohammadne/bookman/library/internal/database/ent"
	"github.com/mohammadne/bookman/library/internal/models"
)

type Book interface {
	GetBook(ctx context.Context, id uint64) (*models.Book, error)
}

func entBookToModelsBook(entCall *ent.Book) *models.Book {
	return nil
}

func (db *database) GetBook(ctx context.Context, id uint64) (*models.Book, error) {
	ctx, span := db.tracer.Start(ctx, "database.book.get")
	defer span.End()

	entBook, err := db.client.Book.Get(ctx, int(id))
	if err != nil {
		if ent.IsNotFound(err) {
			span.RecordError(ErrNotFound)
			return nil, ErrNotFound
		}

		err = fmt.Errorf("error while getting Book from database, error: %w", err)
		span.RecordError(err)
		return nil, err
	}

	return entBookToModelsBook(entBook), nil
}
