package routes

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserById/:userId", func(c *gin.Context) {})
	r.GET("/getUserByEmail/:userEmail")
	r.POST("/user/")
	r.PUT("/user/:userId")
	r.DELETE("/user/:userId")
}
