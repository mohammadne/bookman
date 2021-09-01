package book

import (
	"context"

	"github.com/mohammadne/bookman/library/pkg/database"
	"github.com/mohammadne/go-pkgs/failures"
)

type Repository interface {
	GetByID(ctx context.Context, id uint64) (*Book, failures.Failure)
	GetByTitle(ctx context.Context, title string) (*Book, failures.Failure)
}

type repositoryImpl struct {
	database database.Database
}

func NewRepository(db database.Database) Repository {
	return &repositoryImpl{database: db}
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

	return book, failure
}

func (repository *repositoryImpl) GetByTitle(ctx context.Context, title string) (*Book, failures.Failure) {
	book := new(Book)
	failure := repository.database.Read(
		getById,
		[]interface{}{title},
		book.Id, book.Name, book.Title, book.DateCreated,
	)

	return book, failure
}
