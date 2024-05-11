package method

import (
	"net/http"

	"github.com/Alym62/rest-go/handler/response"
	"github.com/Alym62/rest-go/schemas"
	"github.com/gin-gonic/gin"
)

func FindByIdHandler(context *gin.Context) {
	id := context.Query("id")
	if id == "" {
		response.SendError(context, http.StatusBadRequest, "id is not empty")
		return
	}

	product := schemas.Product{}
	if err := db.First(&product, id).Error; err != nil {
		response.SendError(context, http.StatusNotFound, "id not found")
		return
	}

	response.SendSuccess(context, "find-id-product", product)
}
