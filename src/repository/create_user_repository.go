package repository

import (
	"context"

	"github.com/RenatoHioji/simple-go-api/src/configuration/logger"
	"github.com/RenatoHioji/simple-go-api/src/configuration/rest_err"
	"github.com/RenatoHioji/simple-go-api/src/model"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init creating user")

	query := `INSERT INTO users (email, username, password, age) VALUES ($1, $2, $3, $4) RETURNING ID`

	if _, err := ur.pool.Exec(context.Background(), query, userDomain.GetEmail(), userDomain.GetUsername(), userDomain.GetPassword(), userDomain.GetAge()); err != nil {
		logger.Error("Error creating user", err)
		return nil, rest_err.NewInternalServerErr(err.Error())
	}

	return userDomain, nil
}
