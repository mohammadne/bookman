package users

import "github.com/mohammadne/bookman/user/utils"

type UserDaoImpl struct{}

func (impl UserDaoImpl) Get(userId int64) (*User, *utils.RestError) {
	return nil, nil
}

func (impl UserDaoImpl) Save(user User) *utils.RestError {
	return nil
}
