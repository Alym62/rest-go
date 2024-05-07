package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/ws")
	{
		v1.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Primeira API com Go",
			})
		})
		v1.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "Primeira API com Go",
			})
		})
		v1.POST("/", func(context *gin.Context) {
			context.JSON(http.StatusCreated, gin.H{
				"message": "Primeira API com Go",
			})
		})
		v1.PUT("/", func(context *gin.Context) {
			context.JSON(http.StatusNoContent, gin.H{
				"message": "Primeira API com Go",
			})
		})
		v1.DELETE("/", func(context *gin.Context) {
			context.JSON(http.StatusNoContent, gin.H{
				"message": "Primeira API com Go",
			})
		})
	}
}
