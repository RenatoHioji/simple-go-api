package user

import (
	"github.com/RenatoHioji/simple-go-api/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(service service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: service,
	}
}

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
