package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateHandler(context *gin.Context) {
	context.JSON(http.StatusCreated, gin.H{
		"message": "Primeira API com Go",
	})
}
