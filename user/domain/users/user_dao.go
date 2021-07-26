package users

import "github.com/mohammadne/bookman/user/utils"

type UserDAO interface {
	Get(userId int64) (*User, *utils.RestError)
	Save() *utils.RestError
}

func (user User) Get(userId int64) (*User, *utils.RestError) {
	return nil, nil
}

func (user User) Save() *utils.RestError {
	return nil
}

func measure(g UserDAO) {}

func a() {
	user := User{}
	measure(user)
}
