package database

type Db interface {
	InitDb()
	AddUser(data []interface{}) error
	ListUsers() []interface{}
}
