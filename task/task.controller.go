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

func getUserId(c *gin.Context) int {
	userId := c.MustGet("x-user-id").(int)
	return userId
}

func (c *controller) helloWorld(ctx *gin.Context) {
	ctx.JSON(base.NewApiMessage(http.StatusOK, c.service.HelloWorld()))
}

func (c *controller) getAll(ctx *gin.Context) {
	userId := getUserId(ctx)
	tasks, err := c.service.GetAll(userId)

	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(base.NewApiMessage(http.StatusOK, tasks))
}

func (c *controller) getById(ctx *gin.Context) {
	userId := getUserId(ctx)
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.Error(&Errors.InvalidId)
		return
	}
	task, err := c.service.GetById(id, userId)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(base.NewApiMessage(http.StatusOK, task))
}

func (c *controller) create(ctx *gin.Context) {
	userId := getUserId(ctx)

	var createRequest CreateRequest
	if err := ctx.BindJSON(&createRequest); err != nil {

		ctx.Error(&Errors.InvalidId)
		return
	}

	task, err := c.service.Create(createRequest, userId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(base.NewApiMessage(http.StatusOK, task))
}

func (c *controller) updateById(ctx *gin.Context) {
	userId := getUserId(ctx)

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

	task, err := c.service.UpdateById(id, updateRequest, userId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(base.NewApiMessage(http.StatusOK, task))
}

func (c *controller) deleteById(ctx *gin.Context) {
	userId := getUserId(ctx)

	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.Error(&Errors.InvalidId)
		return
	}

	err = c.service.DeleteById(id, userId)
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
