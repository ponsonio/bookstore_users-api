package users

import (
	"fmt"
	"github.com/jcabrera/bookstore_users-api/utils/errors"
)

var(
	userDB = make(map[int64]*User)
)

func (user *User) Save() *errors.RestError {
	userDB[user.Id] = user
	return nil
}

func (user *User) Get() *errors.RestError {
	result := userDB[user.Id]
	if result == nil{
		return errors.NewNotFoundError(
			fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	user.FirstName = result.FirstName
	user.LastName = result.LastName

	return nil
}