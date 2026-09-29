package service

import (
	"github.com/RenatoHioji/simple-go-api/src/configuration/logger"
	"github.com/RenatoHioji/simple-go-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (*userDomainService) DeleteUser(string) *rest_err.RestErr {
	logger.Info("Init deleting user", zap.String("journey", "delete_user"))

	return nil
}
