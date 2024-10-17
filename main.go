package main

import (
	"go-todo-app/auth"
	"go-todo-app/base"
	"go-todo-app/task"
	"go-todo-app/user"
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
	ginRouter.Use(base.ErrorHandler)
	task.NewRouter(ginRouter).Init()
	user.NewRouter(ginRouter).Init()
	auth.NewRouter(ginRouter).Init()

	ginRouter.Run("localhost:3000")
}
