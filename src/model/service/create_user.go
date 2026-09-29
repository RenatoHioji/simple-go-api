package service

import (
	"fmt"

	"github.com/RenatoHioji/simple-go-api/src/configuration/logger"
	"github.com/RenatoHioji/simple-go-api/src/configuration/rest_err"
	"github.com/RenatoHioji/simple-go-api/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {

	logger.Info("Init creating user", zap.String("journey", "create_user"))

	userDomain.EncryptPassword()

	fmt.Println(userDomain)

	return nil
}
