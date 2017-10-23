package service

import (
	"entity"
	"errors"
)

var (
	// ErrUsernameIsEmpty is the error for empty username
	ErrUsernameIsEmpty = errors.New("username should not be empty")
	// ErrUsernameExists is the error for username already existing
	ErrUsernameExists = errors.New("username already exists")
)

func validateNewUser(user *entity.User) error {
	if len(user.Username) == 0 {
		return ErrUsernameIsEmpty
	}
	if entity.UserModel.FindByUsername(user.Username) != nil {
		return ErrUsernameExists
	}
	return nil
}

// Register registers a user
func Register(username string, password string, email string, phone string) (err error) {
	newUser := &entity.User{
		Username: username,
		Password: password,
		Email:    email,
		Phone:    phone,
	}
	err = validateNewUser(newUser)
	if err != nil {
		return
	}
	entity.UserModel.AddUser(newUser)
	return
}
