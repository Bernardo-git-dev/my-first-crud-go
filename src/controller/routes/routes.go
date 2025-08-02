package routes

import (
	"github.com/Bernardo-git-dev/my-first-crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserByEmail/:email", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.GET("/getUserById/:id", controller.FindUserByID)
	r.PUT("/updateUser/:id", controller.UpdateUser)

	//teste de ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "pong",
		})
	})
}
