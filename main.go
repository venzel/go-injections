package main

import (
	"venzel.com.br/interface/database"
	"venzel.com.br/interface/domain/user"
)

func main() {
	db := &database.Gorm{}
	db.InitDb()

	repository := &user.RepositoryImpl{
		Db: db,
	}

	service := &user.ServiceImpl{
		Repository: repository,
	}

	_, err := service.Create("Marcos", 15)

	if err != nil {
		panic(err)
	}

	result := service.List()

	for _, user := range result {
		println(user.Name)
	}
}
