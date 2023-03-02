package main

import (
	"api/config"
	"api/routes"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	config.Connect()
	routes.Router()
}
