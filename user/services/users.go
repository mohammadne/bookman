package services

import (
	"github.com/mohammadne/bookman/user/domain/users"
	"github.com/mohammadne/bookman/user/utils"
)

func CreateUser(user users.User) (*users.User, *utils.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	return nil, nil

}
