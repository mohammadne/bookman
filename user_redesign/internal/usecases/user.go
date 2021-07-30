package usecases

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/mohammadne/bookman/user/internal"
	"github.com/mohammadne/bookman/user/internal/entities"
	"github.com/mohammadne/bookman/user/pkg/errors"
	"github.com/mohammadne/bookman/user/pkg/logger"
)

type userUsecase struct {
	repo   internal.IUserRepository
	logger logger.Logger
}

// cfg *config.Config , logger logger.Logger
// return &userUsecase{cfg: cfg, commRepo: repo, logger: logger}
// Comments UseCase constructor
func NewUserUsecase(repo internal.IUserRepository) internal.IUserRepository {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) Create(ctx context.Context, user *entities.User) (*entities.User, *errors.RestError) {
	return nil, nil
}

func (u *userUsecase) GetByID(ctx context.Context, id uuid.UUID) (*entities.User, *errors.RestError) {
	return nil, nil
}
