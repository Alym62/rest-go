package router

import (
	"github.com/Alym62/rest-go/handler"
	"github.com/Alym62/rest-go/handler/method"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// TODO: Initializar request
	handler.InitializerHandler()

	v1 := router.Group("/api/v1/ws")
	{
		v1.GET("/", method.FindHandler)
		v1.GET("/id", method.FindByIdHandler)
		v1.POST("/create", method.CreateHandler)
		v1.PUT("/update", method.UpdateHandler)
		v1.DELETE("/delete", method.DeleteHandler)
	}
}
