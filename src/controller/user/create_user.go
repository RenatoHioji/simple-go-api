package user

import (
	"net/http"

	"github.com/RenatoHioji/simple-go-api/src/configuration/logger"
	"github.com/RenatoHioji/simple-go-api/src/configuration/validation"
	"github.com/RenatoHioji/simple-go-api/src/controller/user/requests"
	"github.com/RenatoHioji/simple-go-api/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
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

	domain := model.NewUserDomain(user_request.Email, user_request.Password, user_request.Username, user_request.Age)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("Sucessful creating user", zap.String("journey", "create_user"))

	c.String(http.StatusOK, "")
}
