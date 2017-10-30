package entity_test

import (
	"go/build"
	"runtime"
	"testing"

	"entity"
	"service"
)

func init() {
	if runtime.GOOS == "windows" {
		entity.UserModel.Init(build.Default.GOPATH + "\\tmp\\test_user_data.json")
	} else {
		entity.UserModel.Init("/tmp/test_user_data.json")
	}
}

func TestUserModel(t *testing.T) {
	model := entity.UserModel
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
			t.Errorf("Expect user %d to have username '%s', but got '%s'", index, "test", user.Username)
		}
	}
}

func TestUserService(t *testing.T) {
	model := entity.UserModel
	var err error
	err = service.Register("testUsername", "testPassword", "email@email.com", "12345678912")
	if err != nil {
		t.Fatal(err)
	}
	newUser := model.FindByUsername("testUsername")
	if newUser == nil {
		t.Fatalf(`Expect user '%s' could be found`, "testUsername")
	}
}
