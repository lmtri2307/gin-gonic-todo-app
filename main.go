package main

import (
	"go-todo-app/task"

	"github.com/gin-gonic/gin"
)

func main() {
	ginRouter := gin.Default()
	task.Router.Init(ginRouter)

	ginRouter.Run("localhost:3000")
}
