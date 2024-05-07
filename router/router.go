package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// TODO: Initializer Router
	router := gin.Default()

	// TODO: Initializar Method Routes
	initializeRoutes(router)

	// TODO: Run the server
	router.Run(":8080")
}
