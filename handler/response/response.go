package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendError(context *gin.Context, code int, message string) {
	context.JSON(code, gin.H{
		"message":   message,
		"errorCode": code,
	})
}

func SendSuccess(context *gin.Context, operation string, data interface{}) {
	context.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("operation %s success", operation),
		"data":    data,
	})
}
