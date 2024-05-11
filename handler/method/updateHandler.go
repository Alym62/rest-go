package method

import (
	"net/http"

	"github.com/Alym62/rest-go/handler/request"
	"github.com/Alym62/rest-go/handler/response"
	"github.com/Alym62/rest-go/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateHandler(context *gin.Context) {
	request := request.UpdateProductRequest{}

	// TODO: Validate rerquest
	context.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errf("validation error: %v", err)
		response.SendError(context, http.StatusBadRequest, err.Error())
		return
	}

	// TODO: Validation id is not empty
	id := context.Query("id")
	if id == "" {
		response.SendError(context, http.StatusBadRequest, "id is not empty")
		return
	}

	// TODO: Search id and check not found
	product := schemas.Product{}
	if err := db.First(&product, id).Error; err != nil {
		response.SendError(context, http.StatusNotFound, "id is not found")
		return
	}

	// TODO: Update product
	if request.Name != "" {
		product.Name = request.Name
	}

	if request.Company != "" {
		product.Company = request.Company
	}

	if request.Location != "" {
		product.Location = request.Location
	}

	if request.Price <= 0 {
		product.Price = request.Price
	}

	if request.Description != "" {
		product.Description = request.Description
	}

	// TODO: Save product
	if err := db.Save(&product).Error; err != nil {
		logger.Errf("error updating product: %v", err.Error())
		response.SendError(context, http.StatusInternalServerError, "error updating product")
		return
	}

	response.SendSuccess(context, "update-product", product)
}
