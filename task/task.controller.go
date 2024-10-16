package task

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service *service
}

func (c *controller) helloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, c.service.HelloWorld())
}

func (c *controller) getAll(ctx *gin.Context) {
	tasks := c.service.GetAll()
	ctx.JSON(http.StatusOK, tasks)
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
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (c *controller) create(ctx *gin.Context) {
	var createRequest CreateRequest
	if err := ctx.BindJSON(&createRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.New("invalid task payload").Error()})
		return
	}

	task, err := c.service.Create(createRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errors.New("internal server error"))
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func newController() *controller {
	service := NewService()
	controller := controller{service}

	return &controller
}
