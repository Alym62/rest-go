package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Primeira API com Go",
	})
}
