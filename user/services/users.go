package services

import "github.com/mohammadne/bookman/user/domain"

func CreateUser(user domain.User) (*domain.User, error) {
	return &user, nil
}
