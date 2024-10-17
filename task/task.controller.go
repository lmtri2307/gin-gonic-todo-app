package task

import (
	"go-todo-app/base"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service *service
}

func (c *controller) helloWorld(ctx *gin.Context) {
	ctx.JSON(base.NewApiMessage(http.StatusOK, c.service.HelloWorld()))
}

func (c *controller) getAll(ctx *gin.Context) {
	tasks, err := c.service.GetAll()

	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(base.NewApiMessage(http.StatusOK, tasks))
}

func (c *controller) getById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.Error(&Errors.InvalidId)
		return
	}
	task, err := c.service.GetById(id)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(base.NewApiMessage(http.StatusOK, task))
}

func (c *controller) create(ctx *gin.Context) {
	var createRequest CreateRequest
	if err := ctx.BindJSON(&createRequest); err != nil {

		ctx.Error(&Errors.InvalidId)
		return
	}

	task, err := c.service.Create(createRequest)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(base.NewApiMessage(http.StatusOK, task))
}

func (c *controller) updateById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.Error(&Errors.InvalidId)
		return
	}

	var updateRequest UpdateRequest
	if err := ctx.BindJSON(&updateRequest); err != nil {
		ctx.Error(&Errors.InvalidUpdatePayload)
		return
	}

	task, err := c.service.UpdateById(id, updateRequest)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(base.NewApiMessage(http.StatusOK, task))
}

func (c *controller) deleteById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.Error(&Errors.InvalidId)
		return
	}

	err = c.service.DeleteById(id)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(base.NewApiMessage(http.StatusOK, true))

}

func newController() *controller {
	service := NewService()
	controller := controller{service}

	return &controller
}
