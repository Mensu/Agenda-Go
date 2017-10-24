package service

import (
	"entity"
	"fmt"
	"utils"
)

func validateNewUser(user *entity.User) error {
	if len(user.Username) == 0 {
		return fmt.Errorf("username should not be empty")
	}
	if entity.UserModel.FindByUsername(user.Username) != nil {
		return fmt.Errorf("username '%s' already exists", user.Username)
	}
	// ...TODO
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
	newUser.Password = utils.MD5(newUser.Password)
	entity.UserModel.AddUser(newUser)
	return
}
