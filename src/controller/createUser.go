package controller

import (
	"net/http"

	"github.com/Bernardo-git-dev/my-first-crud-go/src/configuration/logger"
	"github.com/Bernardo-git-dev/my-first-crud-go/src/configuration/validation"
	"github.com/Bernardo-git-dev/my-first-crud-go/src/controller/model/request"
	"github.com/Bernardo-git-dev/my-first-crud-go/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
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

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	logger.Info("User created successfully",
		zap.String("jurney", "CreateUser"),
	)

	c.JSON(http.StatusOK, response)
}
