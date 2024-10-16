package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service *service
}

func (c *controller) helloWorld(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, c.service.HelloWorld())
}

func (c *controller) getAll(ctx *gin.Context) {
	tasks := c.service.GetAll()
	ctx.IndentedJSON(http.StatusOK, tasks)
}

func newController() *controller {
	service := NewService()
	controller := controller{service}

	return &controller
}
