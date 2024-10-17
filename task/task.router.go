package task

import (
	"go-todo-app/auth"

	"github.com/gin-gonic/gin"
)

type router struct {
	authMiddleware *auth.Middleware
	controller     *controller
	engine         *gin.Engine
}

func (r *router) Init() {
	group := r.engine.Group("/tasks")
	group.Use(r.authMiddleware.JwtAuthMiddleware)
	group.GET("/hello-world", r.controller.helloWorld)
	group.GET("/", r.controller.getAll)
	group.GET("/:id", r.controller.getById)
	group.POST("/", r.controller.create)
	group.PUT("/:id", r.controller.updateById)
	group.DELETE("/:id", r.controller.deleteById)
}

func NewRouter(e *gin.Engine) *router {
	router := router{
		authMiddleware: auth.NewMiddleware(),
		controller:     newController(),
		engine:         e,
	}

	return &router
}
