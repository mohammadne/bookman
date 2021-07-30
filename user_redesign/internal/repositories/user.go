package repositories

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mohammadne/bookman/user/internal"
	"github.com/mohammadne/bookman/user/internal/entities"
	"github.com/mohammadne/bookman/user/pkg/errors"
)

type userRepository struct {
	db *sqlx.DB
}

var (
	userDB = make(map[int64]*entities.User)
)

func NewUserRepository(db *sqlx.DB) internal.IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Get(userId int64) (*entities.User, *errors.RestError) {
	result := userDB[userId]
	if result == nil {
		errStr := fmt.Sprintf("user %d not found", userId)
		return nil, errors.NewNotFoundError(errStr)
	}

	return result, nil
}

func (r *userRepository) Save(user *entities.User) *errors.RestError {
	if current := userDB[user.Id]; current != nil {
		if current.Email == user.Email {
			errStr := fmt.Sprintf("email %s already registered", user.Email)
			return errors.NewBadRequestError(errStr)
		}

		errStr := fmt.Sprintf("user %d already exists", user.Id)
		return errors.NewBadRequestError(errStr)
	}

	now := time.Now().UTC()
	user.DateCreated = now.Format("2006-01-02T15:04:05Z")

	userDB[user.Id] = user
	return nil
}
