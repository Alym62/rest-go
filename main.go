package main

import (
	"github.com/Alym62/rest-go/config"
	"github.com/Alym62/rest-go/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initializer Configs
	err := config.Init()
	if err != nil {
		logger.Errf("Error: %v", err)
		return
	}

	// Initializer Router
	router.Initialize()
}
