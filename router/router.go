package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter() {
	r = gin.Default()
	r.GET("/helloWorld", HelloWorld)
}

func Start(addr string) error {
	return r.Run(addr)
}

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!!",
	})
	return
}
