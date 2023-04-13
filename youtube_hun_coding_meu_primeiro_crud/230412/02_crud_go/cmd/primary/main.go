package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello World from the primary executable")
	err := godotenv.Load("secret/.env")
	if err != nil {
		log.Fatal("error loading .env file")
	}
	fmt.Println(os.Getenv("TEST"))

}
