package response

import (
	"fmt"

	"github.com/Alym62/rest-go/schemas"
	"github.com/gin-gonic/gin"
)

func SendError(context *gin.Context, code int, message string) {
	context.JSON(code, gin.H{
		"message":   message,
		"errorCode": code,
	})
}

func SendSuccess(context *gin.Context, code int, operation string, data interface{}) {
	context.JSON(code, gin.H{
		"message": fmt.Sprintf("operation %s success", operation),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateProductResponseSwagger struct {
	Message string                  `json:"message"`
	Data    schemas.ProductResponse `json:"data"`
}

type DeleteProductResponseSwagger struct {
	Message string                  `json:"message"`
	Data    schemas.ProductResponse `json:"data"`
}

type FindByIdProductResponseSwagger struct {
	Message string                  `json:"message"`
	Data    schemas.ProductResponse `json:"data"`
}

type FindProductResponseSwagger struct {
	Message string                  `json:"message"`
	Data    schemas.ProductResponse `json:"data"`
}

type UpdateProductResponseSwagger struct {
	Message string                    `json:"message"`
	Data    []schemas.ProductResponse `json:"data"`
}
