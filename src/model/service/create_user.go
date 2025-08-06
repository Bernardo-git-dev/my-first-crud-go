package service

import (
	"fmt"

	"github.com/Bernardo-git-dev/my-first-crud-go/src/configuration/logger"
	"github.com/Bernardo-git-dev/my-first-crud-go/src/configuration/rest_err"
	"github.com/Bernardo-git-dev/my-first-crud-go/src/model"
	"go.uber.org/zap"
)

var ()

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init CreateUser model", zap.String("jurney", "CreateUser"))

	userDomain.EncryptPassword()
	fmt.Println(userDomain.GetPassword())

	return nil
}
