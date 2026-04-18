package main

import (
	"log"
	"os"
	"relay-engine/internal/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := router.Setup()

	port := os.Getenv("PORT")
	err = r.Run(":" + port)
	if err != nil {
		return
	}
}
