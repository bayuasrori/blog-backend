package main

import (
	"log"

	"goblog/handler"
	"goblog/models"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	models.Migrate()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := handler.Routes()
	r.Run()
}
