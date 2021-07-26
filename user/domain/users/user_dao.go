package users

import "github.com/mohammadne/bookman/user/utils"

type UserDAO interface {
	Get(userId int64) (*User, *utils.RestError)
	Save() *utils.RestError
}
