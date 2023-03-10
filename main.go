package main

import (
	"api/config"
	"api/routes"
	"api/utils"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	utils.LogPort()
	config.Connect()
	routes.Router(os.Getenv("PORT"))
}
