package main

import (
	"fmt"
	"log"
	"os"

	"microservice-go/api"
	"microservice-go/pkg/database"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := api.SetupRouter()
	fmt.Println("Server running on port:", port)
	r.Run(":" + port)
}