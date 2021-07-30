package usecases

import (
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
func NewUserUsecase(repo internal.IUserRepository) internal.IUserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) Create(user *entities.User) (*entities.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := u.repo.Save(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) Get(id int64) (*entities.User, *errors.RestError) {
	return nil, nil
}
