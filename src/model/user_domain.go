package model

func (ud *userDomain) SetId(id string) {
	ud.ID = id
}

type userDomain struct {
	ID       string
	email    string
	password string
	username string
	age      int8
}

func (user *userDomain) GetEmail() string {
	return user.email
}

func (user *userDomain) GetPassword() string {
	return user.password
}

func (user *userDomain) GetUsername() string {
	return user.username
}

func (user *userDomain) GetAge() int8 {
	return user.age
}
