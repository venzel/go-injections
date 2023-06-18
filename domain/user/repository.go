package user

import (
	"venzel.com.br/interface/database"
)

type Repository interface {
	Create(user *User) error
	List() []*User
}

type RepositoryImpl struct {
	Db database.Db
}

func (r *RepositoryImpl) Create(user *User) error {
	gUser := []interface{}{user.Name, user.Age}

	r.Db.AddUser(gUser)

	return nil
}

func (r *RepositoryImpl) List() []*User {
	result := r.Db.ListUsers()

	var users []*User

	for _, user := range result {
		name := user.([]interface{})[0].(string)
		age := user.([]interface{})[1].(uint8)

		users = append(users, &User{name, age})
	}

	return users
}
