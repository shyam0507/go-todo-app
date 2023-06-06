package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/shyam0507/todo-app/pkg/config"
	"github.com/shyam0507/todo-app/pkg/models"
	"github.com/shyam0507/todo-app/pkg/routes"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	log.Printf("Going to connect with database")
	config.Connect()
	config.DB.AutoMigrate(&models.Todo{})
	config.DB.AutoMigrate(&models.User{})

	r := gin.Default()
	routes.RegisterUserRoutes(r)
	routes.RegisterTodoRoutes(r)
	r.Run(":8081")
	log.Printf("todo service running on port 8081")
}
