package method

import (
	"net/http"

	"github.com/Alym62/rest-go/handler/response"
	"github.com/Alym62/rest-go/schemas"
	"github.com/gin-gonic/gin"
)

func FindHandler(context *gin.Context) {
	products := []schemas.Product{}

	if err := db.Find(&products).Error; err != nil {
		response.SendError(context, http.StatusInternalServerError, "Error listing products")
		return
	}

	response.SendSuccess(context, "list-products", products)
}
