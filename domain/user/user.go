package user

import (
	"errors"
)

type User struct {
	Name string
	Age  uint8
}

func (u *User) NewUser(name string, age uint8) (*User, error) {
	user := &User{Name: name, Age: age}

	valid := u.isValid(user)

	if !valid {
		return nil, errors.New("Invalid user")
	}

	return user, nil
}

func (u *User) isValid(user *User) bool {
	if user.Name == "" {
		return false
	}

	return true
}
