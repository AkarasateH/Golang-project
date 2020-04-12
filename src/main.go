package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)
import routers "Golang-project/src/routers"

func init() {
	// Log error if .env file does not exist
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

func main() {
	PORT := os.Getenv("PORT")
	r := routers.SetupRouter()
	
	r.Run(":" + PORT)
}
