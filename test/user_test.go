package entity_test

import (
	"testing"

	"entity"
)

func TestUserModel(t *testing.T) {
	model := entity.UserModel
	model.Init("/tmp/user_data.json")
	newUser := entity.User{
		Username: "test",
		Password: "test",
		Email:    "test@test.com",
		Phone:    "12345678909",
	}
	model.AddUser(&newUser)
	foundUsers := model.FindBy(func(oneUser *entity.User) bool {
		return oneUser.Username == "test"
	})
	for index, user := range foundUsers {
		if user.Username != "test" {
			t.Errorf(`Expect User %d to have username "%s", but got "%s"`, index, "test", user.Username)
		}
	}
}
