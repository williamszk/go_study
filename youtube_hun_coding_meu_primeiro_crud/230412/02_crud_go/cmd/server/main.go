package main

import (
	"crud_go/config/logger"
	resource_user "crud_go/controller/resources/user"
	"crud_go/controller/routes"
	"crud_go/model/user/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	logger.Info("About to start application")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// should all dependencies be change in main?
	// or they could come from a config file?
	// initialize services (the dependencies)
	userDomainService := service.NewUserDomainService()
	userController := resource_user.NewUserControllerInterface(userDomainService)

	// initialize routes
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err = router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
