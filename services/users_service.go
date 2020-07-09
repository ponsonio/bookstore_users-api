package services

import (
	"github.com/jcabrera/bookstore_users-api/domain/users"
	"github.com/jcabrera/bookstore_users-api/utils/crypto_utils"
	"github.com/jcabrera/bookstore_users-api/utils/date_utils"
	"github.com/jcabrera/bookstore_users-api/utils/errors"
)

var(
	UsersService userServiceInterface =  &userService{}
)

type userService struct {

}

type userServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestError)
	GetUser(int64) (*users.User, *errors.RestError)
	UpdateUser(bool, users.User)(*users.User, *errors.RestError)
	DeleteUser(int64) *errors.RestError
	Search(string) (users.Users, *errors.RestError)
	LoginUser(request users.LoginRequest) (*users.User, *errors.RestError)
}

func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil{
		return nil, err
	}
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDatabaseFormatString()
	user.Password = crypto_utils.GetMd5(user.Password)
	if err:= user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) GetUser(id int64) (*users.User, *errors.RestError) {
	result := &users.User{Id: id}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *userService) UpdateUser(isPartial bool, user users.User)(*users.User, *errors.RestError) {
	current := &users.User{Id: user.Id}
	if err := current.Get(); err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}

	} else {
		current.FirstName = user.FirstName
		current.LastName =user.LastName
		current.Email = user.Email
	}


	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *userService) DeleteUser(userID int64) *errors.RestError {
	user := &users.User{Id: userID}
	return user.Delete()
}

func (s *userService) Search(status string) (users.Users, *errors.RestError) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func  (s *userService) 	LoginUser(request users.LoginRequest) (*users.User, *errors.RestError) {
	dao := &users.User{
		Email: request.Email,
		Password: crypto_utils.GetMd5(request.Password),
	}
	if err:= dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}


