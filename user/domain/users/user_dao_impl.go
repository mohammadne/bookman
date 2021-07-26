package users

import (
	"fmt"

	"github.com/mohammadne/bookman/user/utils"
)

type MySQLUserDao struct{}

var (
	userDB = make(map[int64]*User)
)

func (MySQLUserDao) Get(userId int64) (*User, *utils.RestError) {
	result := userDB[userId]
	if result == nil {
		errStr := fmt.Sprintf("user %d not found", userId)
		return nil, utils.NewNotFoundError(errStr)
	}

	return result, nil
}

func (MySQLUserDao) Save(user *User) *utils.RestError {
	if current := userDB[user.Id]; current != nil {
		if current.Email == user.Email {
			errStr := fmt.Sprintf("email %s already registered", user.Email)
			return utils.NewBadRequestError(errStr)
		}

		errStr := fmt.Sprintf("user %d already exists", user.Id)
		return utils.NewBadRequestError(errStr)
	}

	userDB[user.Id] = user
	return nil
}
