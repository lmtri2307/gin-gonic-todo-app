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
	ctx.JSON(http.StatusOK, c.service.HelloWorld())
}

func (c *controller) getAll(ctx *gin.Context) {
	tasks, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Task Payload"})
		return
	}

	task, err := c.service.Create(createRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (c *controller) updateById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id"})
		return
	}

	var updateRequest UpdateRequest
	if err := ctx.BindJSON(&updateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Update Payload"})
		return
	}

	task, err := c.service.UpdateById(id, updateRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (c *controller) deleteById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id"})
		return
	}

	err = c.service.DeleteById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task Not Found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})

}

func newController() *controller {
	service := NewService()
	controller := controller{service}

	return &controller
}
