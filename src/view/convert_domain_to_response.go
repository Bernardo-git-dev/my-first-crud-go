package view

import (
	"github.com/Bernardo-git-dev/my-first-crud-go/src/controller/model/response"
	"github.com/Bernardo-git-dev/my-first-crud-go/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Name:  userDomain.GetName(),
		Email: userDomain.GetEmail(),
		Age:   userDomain.GetAge(),
	}
}
