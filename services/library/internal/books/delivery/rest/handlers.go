package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/library/internal/books"
)

type Handler interface {
	getBook(ctx echo.Context) error
}

type handler struct {
	usecase books.Usecase
}

func NewHandler(usecase books.Usecase) Handler {
	return &handler{usecase: usecase}
}

func (rest *handler) getBook(ctx echo.Context) error {
	return nil
}
