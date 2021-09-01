package books

import (
	"context"

	"github.com/mohammadne/bookman/library/pkg/database"
	"github.com/mohammadne/bookman/library/pkg/failures"
	"github.com/mohammadne/bookman/library/pkg/logger"
)

type Repository interface {
	GetByID(ctx context.Context, id uint64) (*Book, failures.Failure)
	GetByTitle(ctx context.Context, title string) (*Book, failures.Failure)
}

type repositoryImpl struct {
	database database.Database
	logger   logger.Logger
}

func NewRepository(db database.Database, logger logger.Logger) Repository {
	return &repositoryImpl{database: db, logger: logger}
}

// Queries
const (
	getById    = "SELECT id, name, title, date_created FROM books WHERE id=?;"
	getByTitle = "SELECT id, name, title, date_created FROM books WHERE title=?;"
)

func (repository *repositoryImpl) GetByID(ctx context.Context, id uint64) (*Book, failures.Failure) {
	book := new(Book)
	failure := repository.database.Read(
		getById,
		[]interface{}{id},
		book.Id, book.Name, book.Title, book.DateCreated,
	)

	if failure != nil {
		repository.logger.Error(failure.Message(), logger.Error(failure))
		return nil, failure
	}

	return book, nil
}

func (repository *repositoryImpl) GetByTitle(ctx context.Context, title string) (*Book, failures.Failure) {
	book := new(Book)
	failure := repository.database.Read(
		getByTitle,
		[]interface{}{title},
		book.Id, book.Name, book.Title, book.DateCreated,
	)

	if failure != nil {
		repository.logger.Error(failure.Message(), logger.Error(failure))
		return nil, failure
	}

	return book, nil
}
