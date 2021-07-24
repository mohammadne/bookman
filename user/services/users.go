package services

import "github.com/mohammadne/bookman/user/domain/users"

func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
