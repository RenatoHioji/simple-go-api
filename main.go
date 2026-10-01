package main

import (
	"context"
	"log"

	"github.com/RenatoHioji/simple-go-api/src/configuration/database/postgresql"
	"github.com/RenatoHioji/simple-go-api/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pool, err := postgresql.NewPostgresqlConnection(context.Background())

	if err != nil {
		log.Fatal("Error trying to connect to db", err)
	}

	user_controller := initDependencies(pool)
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, user_controller)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
