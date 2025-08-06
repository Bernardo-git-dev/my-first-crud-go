package routes

import (
	"github.com/Bernardo-git-dev/my-first-crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, UserController controller.UserControllerInterface) {

	r.POST("/users", UserController.CreateUser)
	r.GET("/users/:user_id", UserController.FindUserByID)
	r.GET("/users/email/:email", UserController.FindUserByEmail)
	r.PUT("/users/:user_id", UserController.UpdateUser)
	r.DELETE("/users/:user_id", UserController.DeleteUser)

	//teste de ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "pong",
		})
	})
}
