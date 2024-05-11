package method

import (
	"net/http"

	"github.com/Alym62/rest-go/handler/response"
	"github.com/Alym62/rest-go/schemas"
	"github.com/gin-gonic/gin"
)

//	@BasePath	/api/v1/ws

//	@Summary		Find product
//	@Description	Listing products
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.FindProductResponseSwagger
//	@Failure		500	{object}	response.ErrorResponse
//	@Router			/ [get]
func FindHandler(context *gin.Context) {
	products := []schemas.Product{}

	if err := db.Find(&products).Error; err != nil {
		response.SendError(context, http.StatusInternalServerError, "Error listing products")
		return
	}

	response.SendSuccess(context, http.StatusOK, "list-products", products)
}
