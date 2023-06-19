package database

type Gorm struct {
	Users []User
}

type User struct {
	Name string
	Age  uint8
}

func (g *Gorm) InitDb() {
	tiago := User{Name: "Tiago", Age: 15}
	cintia := User{Name: "Cintia", Age: 20}

	g.Users = append(g.Users, tiago, cintia)
}

func (g *Gorm) AddUser(data []interface{}) error {
	name := data[0]
	age := data[1]

	user := User{Name: name.(string), Age: age.(uint8)}

	g.Users = append(g.Users, user)

	return nil
}

func (g *Gorm) ListUsers() []interface{} {
	users := []interface{}{}

	for _, user := range g.Users {
		arr := []interface{}{user.Name, user.Age}
		users = append(users, arr)
	}

	return users
}
