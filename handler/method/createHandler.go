package method

import (
	"net/http"

	"github.com/Alym62/rest-go/handler/request"
	"github.com/Alym62/rest-go/handler/response"
	"github.com/Alym62/rest-go/schemas"
	"github.com/gin-gonic/gin"
)

func CreateHandler(context *gin.Context) {
	request := request.CreateProductRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errf("error validation: %v", err.Error())
		response.SendError(context, http.StatusBadRequest, err.Error())
		return
	}

	product := schemas.Product{
		Name:        request.Name,
		Company:     request.Company,
		Location:    request.Location,
		Price:       request.Price,
		Description: request.Description,
	}

	if err := db.Create(&product).Error; err != nil {
		logger.Errf("error creating product: %v", err.Error())
		response.SendError(context, http.StatusInternalServerError, "error creating product on database")
		return
	}

	response.SendSuccess(context, "create-product", product)
}
