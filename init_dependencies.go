package main

import (
	"github.com/RenatoHioji/simple-go-api/src/controller"
	"github.com/RenatoHioji/simple-go-api/src/model/service"
	"github.com/RenatoHioji/simple-go-api/src/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

func initDependencies(pool *pgxpool.Pool) controller.UserControllerInterface {
	repo := repository.NewUserRepository(pool)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
