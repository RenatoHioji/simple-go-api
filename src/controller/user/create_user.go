package user

import (
	"github.com/RenatoHioji/simple-go-api/src/configuration/validation"
	"github.com/RenatoHioji/simple-go-api/src/controller/user/requests"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user_request requests.UserRequest

	if err := c.ShouldBindJSON(&user_request); err != nil {
		rest_err := validation.ValidateUserError(err)

		c.JSON(rest_err.Code, rest_err)
		return
	}

}
