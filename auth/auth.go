package auth

import "fmt"

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type AuthService interface {
	Register(firstName, lastName, email, password string) string
	Login(email, password string) string
	ForgotPass(email, newPassword string) string
}

type Auth struct {
	Users []User
}

func NewAuth() *Auth {
	return &Auth{
		Users: []User{},
	}
}

func recoverHandler() {
	if r := recover(); r != nil {
		fmt.Println("Error:", r)
	}
}
