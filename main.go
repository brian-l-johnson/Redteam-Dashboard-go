package main

import (
	"log"

	"github.com/brian-l-johnson/Redteam-Dashboard-go/v2/models"
	"github.com/brian-l-johnson/Redteam-Dashboard-go/v2/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to log .env")
	}

	models.Init()
	server.Init()

}
