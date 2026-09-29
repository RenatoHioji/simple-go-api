package service

import (
	"github.com/RenatoHioji/simple-go-api/src/configuration/logger"
	"github.com/RenatoHioji/simple-go-api/src/configuration/rest_err"
	"github.com/RenatoHioji/simple-go-api/src/model"
	"go.uber.org/zap"
)

func (*userDomainService) UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init deleting user", zap.String("journey", "update_user"))

	return nil
}
