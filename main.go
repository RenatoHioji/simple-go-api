package main

import (
	"log"

	"github.com/RenatoHioji/simple-go-api/src/configuration/database/postgresql"
	"github.com/RenatoHioji/simple-go-api/src/controller/routes"
	"github.com/RenatoHioji/simple-go-api/src/controller/user"
	"github.com/RenatoHioji/simple-go-api/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	postgresql.InitPostgresql()

	user_service := service.NewUserDomainService()
	user_controller := user.NewUserControllerInterface(user_service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, user_controller)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
