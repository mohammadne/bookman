package database_impl

import (
	"context"

	"github.com/mohammadne/bookman/library/internal/models"
)

type author interface {
	GetAuthor(ctx context.Context, id uint64) *models.Author
}

func (db *database) GetAuthor(ctx context.Context, id uint64) *models.Author {
	return nil
}
