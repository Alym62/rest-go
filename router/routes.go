package router

import (
	docs "github.com/Alym62/rest-go/docs"
	"github.com/Alym62/rest-go/handler/method"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// TODO: Initializar request
	method.InitializerHandlerLogger()

	basePath := "/api/v1/ws"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.GET("/", method.FindHandler)
		v1.GET("/find", method.FindByIdHandler)
		v1.POST("/create", method.CreateHandler)
		v1.PUT("/update", method.UpdateHandler)
		v1.DELETE("/delete", method.DeleteHandler)
	}

	// TODO: Initializer swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
