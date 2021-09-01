package book

import (
	"context"

	"github.com/mohammadne/go-pkgs/failures"
)

type Usecase interface {
	GetById(context.Context, uint64) (*Book, failures.Failure)
}

type usecaseImpl struct {
	repository Repository
}

func NewUsecase() Usecase {
	return &usecaseImpl{}
}

func (usecase *usecaseImpl) GetById(ctx context.Context, id uint64) (*Book, failures.Failure) {
	return usecase.repository.GetByID(ctx, id)
}
