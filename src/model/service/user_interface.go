package service

import (
	"github.com/RenatoHioji/simple-go-api/src/configuration/rest_err"
	"github.com/RenatoHioji/simple-go-api/src/model"
	"github.com/RenatoHioji/simple-go-api/src/repository"
)

func NewUserDomainService(user_repository repository.UserRepository) UserDomainService {
	return &userDomainService{user_repository: user_repository}
}

type userDomainService struct {
	user_repository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
