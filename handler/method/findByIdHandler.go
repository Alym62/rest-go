package method

import (
	"net/http"

	"github.com/Alym62/rest-go/handler/response"
	"github.com/Alym62/rest-go/schemas"
	"github.com/gin-gonic/gin"
)

//	@BasePath	/api/v1/ws

//	@Summary		FindById product
//	@Description	Find by id product
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			id	query		string	true	"Product identification "
//	@Success		200	{object}	response.FindByIdProductResponseSwagger
//	@Failure		400	{object}	response.ErrorResponse
//	@Failure		404	{object}	response.ErrorResponse
//	@Router			/find [get]
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

	response.SendSuccess(context, http.StatusOK, "find-id-product", product)
}
