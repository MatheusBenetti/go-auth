package main

import (
	"log"
	"os"

	"github.com/MatheusBenetti/go-auth/models"
	"github.com/MatheusBenetti/go-auth/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := models.Database{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	models.ConnectDB(config)

	routes.AuthRoutes(r)

	r.Run(":8080")
}
