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

func FindByIdHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Primeira API com Go",
	})
}

func CreateHandler(context *gin.Context) {
	context.JSON(http.StatusCreated, gin.H{
		"message": "Primeira API com Go",
	})
}

func UpdateHandler(context *gin.Context) {
	context.JSON(http.StatusNoContent, gin.H{
		"message": "Primeira API com Go",
	})
}

func DeleteHandler(context *gin.Context) {
	context.JSON(http.StatusNoContent, gin.H{
		"message": "Primeira API com Go",
	})
}
