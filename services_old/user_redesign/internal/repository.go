package internal

import (
	"github.com/mohammadne/bookman/user/internal/entities"
	"github.com/mohammadne/bookman/user/pkg/errors"
)

type IUserRepository interface {
	Save(user *entities.User) *errors.RestError
	Get(id int64) (*entities.User, *errors.RestError)
}
