package entity_test

import (
	"go/build"
	"runtime"
	"service"
	"testing"

	"entity"
)

func init() {
	if runtime.GOOS == "windows" {
		entity.CurSessionModel.Init(build.Default.GOPATH + "\\tmp\\test_curUser.json")
	} else {
		entity.CurSessionModel.Init("/tmp/test_curUser.json")
	}
}

func TestSessionModel(t *testing.T) {
	model := entity.CurSessionModel
	model.SetCurUser(&entity.User{
		Username: "test",
	})
	if model.GetCurUser() != "test" {
		t.Errorf(`Expect current user to be '%s'`, "test")
	}
	model.SetCurUser(&entity.User{})
}

func TestSessionService(t *testing.T) {
	err := service.Register("sessionUsername", "testPassword", "email@email.com", "12345678912")
	if err != nil {
		t.Fatal(err)
	}
	if err := service.Login("sessionUsername", "wrongpassword"); err != service.ErrInvalidCredentials {
		t.Fatal(err)
	}
	if err := service.Login("sessionUsername", "testPassword"); err != nil {
		t.Fatal(err)
	}
	if err := service.Logout(); err != nil {
		t.Fatal(err)
	}

	if err := service.Login("sessionUsername", "testPassword"); err != nil {
		t.Fatal(err)
	}
	err = service.DeleteUser()
}
