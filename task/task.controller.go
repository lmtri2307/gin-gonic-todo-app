package task

import (
	"net/http"
	"strconv"

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

func (c *controller) getById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	task, err := c.service.GetById(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, task)
}

func newController() *controller {
	service := NewService()
	controller := controller{service}

	return &controller
}
