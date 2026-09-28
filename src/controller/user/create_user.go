package user

import (
	"github.com/RenatoHioji/simple-go-api/src/configuration/logger"
	"github.com/RenatoHioji/simple-go-api/src/configuration/validation"
	"github.com/RenatoHioji/simple-go-api/src/controller/user/requests"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	var user_request requests.UserRequest

	logger.Info("init", zap.String("journey", "create_user"))

	if err := c.ShouldBindJSON(&user_request); err != nil {
		rest_err := validation.ValidateUserError(err)
		logger.Error("error creating user", err, zap.String("journey", "create_user"))
		c.JSON(rest_err.Code, rest_err)
		return
	}

	logger.Info("Sucessful creating user", zap.String("journey", "create_user"))
}
