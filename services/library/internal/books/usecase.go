package books

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

func NewUsecase(repository Repository) Usecase {
	return &usecaseImpl{repository: repository}
}

func (usecase *usecaseImpl) GetById(ctx context.Context, id uint64) (*Book, failures.Failure) {
	return usecase.repository.GetByID(ctx, id)
}
