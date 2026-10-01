package routes

import (
	"github.com/RenatoHioji/simple-go-api/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, user_controller controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", user_controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", user_controller.FindUserByEmail)
	r.POST("/user/", user_controller.CreateUser)
	r.PUT("/user/:userId", user_controller.UpdateUser)
	r.DELETE("/user/:userId", user_controller.DeleteUser)
}
