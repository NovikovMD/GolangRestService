package handler

import (
	"awesomeGoProject/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RouteHandler struct {
	Services *service.Service
}

func NewHandler(srv *service.Service) *RouteHandler {
	return &RouteHandler{Services: srv}
}

func (i *RouteHandler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("")
	{
		auth.GET("/hello", i.handle)
	}

	return router
}

func (i *RouteHandler) handle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
