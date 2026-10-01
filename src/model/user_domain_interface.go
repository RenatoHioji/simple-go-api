package model

import "github.com/RenatoHioji/simple-go-api/src/configuration/rest_err"

type UserDomainInterface interface {
	SetId(string)
	GetEmail() string
	GetPassword() string
	GetUsername() string
	GetAge() int8
	EncryptPassword() *rest_err.RestErr
}

func NewUserDomain(
	email, password, username string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		username: username,
		age:      age,
	}
}
