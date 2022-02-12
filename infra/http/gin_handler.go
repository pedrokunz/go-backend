package http

import "github.com/gin-gonic/gin"

func NewGinHandler() *gin.Engine {
	r := gin.Default()
	return r
}
