package model

import (
	"github.com/RenatoHioji/simple-go-api/src/configuration/logger"
	"github.com/RenatoHioji/simple-go-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (*UserDomain) UpdateUser(string) *rest_err.RestErr {
	logger.Info("Init deleting user", zap.String("journey", "update_user"))

	return nil
}
