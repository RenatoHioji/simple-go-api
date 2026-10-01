package view

import (
	"github.com/RenatoHioji/simple-go-api/src/controller/response"
	"github.com/RenatoHioji/simple-go-api/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:       "1",
		Email:    userDomain.GetEmail(),
		Username: userDomain.GetUsername(),
		Age:      userDomain.GetAge(),
	}
}
