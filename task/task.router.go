package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type router struct {
}

func (r *router) Init(e *gin.Engine) {
	group := e.Group("/tasks")
	group.GET("/hello-world", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Hello world"})
	})
}

var Router router = router{}
