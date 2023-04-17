package main

import (
	"crud_go/controller/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("secret/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err = router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
