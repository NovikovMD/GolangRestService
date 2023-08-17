package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Initializer struct {
}

func (i *Initializer) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("")
	{
		auth.GET("/hello", i.handle)
	}

	return router
}

func (i *Initializer) handle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
