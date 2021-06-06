package main

import (
	"fmt"
	"os"

	"order-food-app-golang/model"
	"order-food-app-golang/server"

	"github.com/joho/godotenv"
)

func main() {
	// if os.Getenv("ENV") == "development" {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error getting env, not comming through %v", err)
	}
	// }
	port := os.Getenv("PORT")
	if port == "" {
		port = "1213"
	}

	// DB Init
	model.Init()

	// Start route
	r := server.Init()
	r.Run(":" + port)
}
