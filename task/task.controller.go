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

func newController() *controller {
	service := newService()
	controller := controller{service}

	return &controller
}
