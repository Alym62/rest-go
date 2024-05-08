package method

import (
	"github.com/gin-gonic/gin"
)

func CreateHandler(context *gin.Context) {
	request := struct{}{}

	context.BindJSON(&request)
}
