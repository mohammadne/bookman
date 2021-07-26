package users

import (
	"strings"

	"github.com/mohammadne/bookman/user/utils"
)

func (user *User) Validate() *utils.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return utils.NewBadRequestError("invalid email address")
	}

	return nil
}
