package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	database "Golang-project/src/config"
	"Golang-project/src/routers"
)

func init() {
	// Log error if .env file does not exist
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

func main() {
	PORT := os.Getenv("PORT")

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")

	if err != nil {
		fmt.Println("String Convert Error: ", err)
	}

	dbConfig := database.Config{
		User: dbUser,
		Host: dbHost,
		Port: dbPort,
		Name: dbName,
	}

	db, dbErr := database.Open(dbConfig)
	defer db.Close()
	fmt.Printf("DB Connected on %s:%d \n", dbHost, dbPort)

	if dbErr != nil {
		fmt.Println("DB Connection Error: ", dbErr)
	}

	r := routers.SetupRouter()

	r.Run(":" + PORT)
}
