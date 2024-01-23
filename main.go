package main

import (
	"github.com/claudineyveloso/gopportunities.git/config"
	"github.com/claudineyveloso/gopportunities.git/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	// Initialize configs
	err := config.Load()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		panic(err)
	}
	router.Initialize()
}
