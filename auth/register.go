package auth

import (
	"crypto/md5"
	"encoding/hex"
)

func hash(pw string) string {
	hashPw := md5.Sum([]byte(pw))
	return hex.EncodeToString(hashPw[:])
}

func (a *Auth) Register(firstName, lastName, email, password, confirmPass string) string {

	for _, user := range a.Users {
		if user.Email == email {
			panic("Email already registered")
		}
	}

	if password != confirmPass {
		panic("Password not match")
	}

	newUser := User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  hash(password),
	}

	a.Users = append(a.Users, newUser)
	return "Registration successful"
}
