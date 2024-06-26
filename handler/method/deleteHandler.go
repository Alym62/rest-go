package method

import (
	"fmt"
	"net/http"

	"github.com/Alym62/rest-go/handler/response"
	"github.com/Alym62/rest-go/schemas"
	"github.com/gin-gonic/gin"
)

//	@BasePath	/api/v1/ws

//	@Summary		Delete product
//	@Description	Delete a   product
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			id	query		string	true	"Product identification "
//	@Success		204	{object}	response.DeleteProductResponseSwagger
//	@Failure		400	{object}	response.ErrorResponse
//	@Failure		404	{object}	response.ErrorResponse
//	@Failure		500	{object}	response.ErrorResponse
//	@Router			/delete [delete]
func DeleteHandler(context *gin.Context) {
	id := context.Query("id")
	if id == "" {
		response.SendError(context, http.StatusBadRequest, "id is not empyt")
		return
	}

	product := schemas.Product{}

	// TODO: Find product
	if err := db.First(&product, id).Error; err != nil {
		response.SendError(context, http.StatusNotFound, fmt.Sprintf("product with id: %s not found", id))
		return
	}

	// TODO: Delete product
	if err := db.Delete(&product).Error; err != nil {
		response.SendError(context, http.StatusInternalServerError, fmt.Sprintf("error deleting product with id: %s", id))
		return
	}

	response.SendSuccess(context, http.StatusNoContent, "delete-product", product)
}
