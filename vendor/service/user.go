package service

import (
	"entity"
	"fmt"
	"utils"
)

type user = entity.User

func validateNewUser(user *user) error {
	if len(user.Username) == 0 {
		return fmt.Errorf("username should not be empty")
	}
	if entity.UserModel.FindByUsername(user.Username) != nil {
		return fmt.Errorf("username '%s' already exists", user.Username)
	}
	if len(user.Password) == 0 {
		return fmt.Errorf("password should not be empty")
	}
	// ...TODO
	return nil
}

// Register registers a user
func Register(username string, password string, email string, phone string) (err error) {
	newUser := &user{
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

// FindAll find all registered users
func FindAll() ([]user, error) {
	curUserName := entity.CurSessionModel.GetCurUser()
	if len(curUserName) == 0 {
		return nil, fmt.Errorf("You are not logined")
	}
	users := entity.UserModel.FindBy(func(u *user) bool {
		return true
	})
	return users, nil
}

// DeleteUser delete the user and meetings belong to him and he participates
func DeleteUser() error {
	if err := checkIfLoggedin(); err != nil {
		return err
	}

	curUserName := entity.CurSessionModel.GetCurUser()
	meetings := entity.MeetingModel.FindBy(func(m *meeting) bool {
		// check from speecher field
		if m.Speecher == curUserName {
			return true
		}
		// check from participator field
		for _, participator := range m.Participators {
			if participator == curUserName {
				return true
			}
		}
		return false
	})
	for _, meeting := range meetings {
		err := DeleteFromMeeting(meeting.Title)
		if err != nil {
			return err
		}
	}
	Logout()
	entity.UserModel.DeleteUser(entity.UserModel.FindByUsername(curUserName))
	return nil
}

// IsPwdMatch checks whether provided password matches that of the user
func IsPwdMatch(user *user, password string) bool {
	return utils.MD5(password) == user.Password
}
