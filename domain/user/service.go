package user

type Service interface {
	Create(name string, age uint8) (*User, error)
	List() []*User
}

type ServiceImpl struct {
	Repository Repository
}

func (s *ServiceImpl) Create(name string, age uint8) (*User, error) {
	var user *User

	user, err := user.NewUser(name, age)

	if err != nil {
		return nil, err
	}

	err = s.Repository.Create(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *ServiceImpl) List() []*User {
	return s.Repository.List()
}
