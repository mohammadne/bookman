package usecase

import "github.com/mohammadne/bookman/library/internal/books/repository"

type usecaseImpl struct {
	repo repository.Repository
	// redisRepo
}

func New(repo repository.Repository) Usecase {
	return &usecaseImpl{repo: repo}
}
