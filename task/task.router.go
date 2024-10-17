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
	group.GET("/", r.controller.getAll)
	group.GET("/:id", r.controller.getById)
	group.POST("/", r.controller.create)
	group.PUT("/:id", r.controller.updateById)
	group.DELETE("/:id", r.controller.deleteById)
}

func NewRouter(e *gin.Engine) *router {
	controller := newController()
	router := router{controller, e}

	return &router
}
