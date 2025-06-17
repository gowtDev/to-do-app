package tests

import (
	"testing"

	"github.com/gowtDev/to-do-app/data"
)

func TestCreateUser(t *testing.T) {
	data.InitDB()
	user := data.User{Name: "David", Email: "david@gmail.com", Password: "password"}
	token, err := user.CreateUser()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if token == "" {
		t.Error("Expected a token, got empty string")
	}
}

func TestLogin(t *testing.T) {
	data.InitDB()
	user := data.User{Email: "david@gmail.com", Password: "password"}
	token, err := user.Login()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if token == "" {
		t.Error("Expected a token, got empty string")
	}
}
