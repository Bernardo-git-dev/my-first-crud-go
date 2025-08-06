package controller

import (
	"net/http"

	"github.com/Bernardo-git-dev/my-first-crud-go/src/configuration/logger"
	"github.com/Bernardo-git-dev/my-first-crud-go/src/configuration/validation"
	"github.com/Bernardo-git-dev/my-first-crud-go/src/controller/model/request"
	"github.com/Bernardo-git-dev/my-first-crud-go/src/model"
	"github.com/Bernardo-git-dev/my-first-crud-go/src/model/service"
	"github.com/Bernardo-git-dev/my-first-crud-go/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("jurney", "CreateUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("jurney", "CreateUser"),
		)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	service := service.NewUserDomainService()
	if err := service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("jurney", "CreateUser"),
	)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}
