package main

import (
	"log"
	"os"
	"taptoeat-be/models"
	"taptoeat-be/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")

	router := gin.New()
	routes.Routes(router)
	models.CreateConnection(dbname, user, pass)

	router.Run(":8800")
}
