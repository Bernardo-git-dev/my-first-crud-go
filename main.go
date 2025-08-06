package main

import (
	"log"

	"github.com/Bernardo-git-dev/my-first-crud-go/src/controller"
	"github.com/Bernardo-git-dev/my-first-crud-go/src/controller/routes"
	"github.com/Bernardo-git-dev/my-first-crud-go/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":3333"); err != nil {
		log.Fatal(err)
	}
}
