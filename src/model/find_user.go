package model

import (
	"github.com/RenatoHioji/simple-go-api/src/configuration/logger"
	"github.com/RenatoHioji/simple-go-api/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (*UserDomain) FindUser(string) (*UserDomain, *rest_err.RestErr) {
	logger.Info("Init search user", zap.String("journey", "find_user"))

	return nil, nil
}
