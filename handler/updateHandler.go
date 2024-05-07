package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateHandler(context *gin.Context) {
	context.JSON(http.StatusNoContent, gin.H{
		"message": "Primeira API com Go",
	})
}
