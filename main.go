package main

import (
	"go-todo-app/task"

	"github.com/gin-gonic/gin"
)

func main() {
	ginRouter := gin.Default()
	task.NewRouter(ginRouter).Init()

	ginRouter.Run("localhost:3000")
}
