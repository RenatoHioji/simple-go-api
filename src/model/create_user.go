package model

import (
	"fmt"

	"github.com/RenatoHioji/simple-go-api/src/configuration/logger"
	"github.com/RenatoHioji/simple-go-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init creating user", zap.String("journey", "create_user"))

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}
