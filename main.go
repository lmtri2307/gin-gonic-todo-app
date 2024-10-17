package main

import (
	"go-todo-app/task"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	ginRouter := gin.Default()
	task.NewRouter(ginRouter).Init()

	ginRouter.Run("localhost:3000")
}
