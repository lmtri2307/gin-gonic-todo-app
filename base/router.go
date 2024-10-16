package base

import "github.com/gin-gonic/gin"

type Router interface {
	Init(e *gin.Engine)
}
