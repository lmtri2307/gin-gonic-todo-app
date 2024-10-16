package task

import (
	"github.com/gin-gonic/gin"
)

type router struct {
	controller *controller
	engine     *gin.Engine
}

func (r *router) Init() {
	group := r.engine.Group("/tasks")
	group.GET("/hello-world", r.controller.helloWorld)
}

func NewRouter(e *gin.Engine) *router {
	controller := newController()
	router := router{controller, e}

	return &router
}
