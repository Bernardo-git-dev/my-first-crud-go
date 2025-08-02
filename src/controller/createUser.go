package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/Bernardo-git-dev/my-first-crud-go/src/configuration/rest_err"
	"github.com/Bernardo-git-dev/my-first-crud-go/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("there are some incorrect filds, error=%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}

