package routes

import (
	"github.com/RenatoHioji/simple-go-api/src/controller/user"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserById/:userId", user.FindUserById)
	r.GET("/getUserByEmail/:userEmail", user.FindUserByEmail)
	r.POST("/user/", user.CreateUser)
	r.PUT("/user/:userId", user.UpdateUser)
	r.DELETE("/user/:userId", user.DeleteUser)
}
