package services

import (
	"github.com/mohammadne/bookman/user/domain"
	"github.com/mohammadne/bookman/user/utils"
)

func CreateUser(user domain.User) (*domain.User, *utils.RestError) {
	return &user, nil
}
