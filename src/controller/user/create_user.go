package user

import (
	"fmt"

	"github.com/RenatoHioji/simple-go-api/src/configuration/rest_err"
	"github.com/RenatoHioji/simple-go-api/src/controller/user/requests"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user_request requests.UserRequest

	if err := c.ShouldBindJSON(&user_request); err != nil {
		rest_err := rest_err.NewBadRequestErr(fmt.Sprintf("There are some incorrect fields, err= %s", err))

		c.JSON(rest_err.Code, rest_err)
		return
	}

}
