package router

import (
	"github.com/Alym62/rest-go/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/ws")
	{
		v1.GET("/", handler.FindHandler)
		v1.GET("/id", handler.FindByIdHandler)
		v1.POST("/create", handler.CreateHandler)
		v1.PUT("/update", handler.UpdateHandler)
		v1.DELETE("/delete", handler.DeleteHandler)
	}
}
