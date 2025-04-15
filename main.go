package main

import (
	"github.com/GarotoCowboy/Goportunities.git/config"
	"github.com/GarotoCowboy/Goportunities.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initialize configs
	err := config.Init()
	if err != nil {
		logger.ErrorF("config initialization error: %v", err)
		return
	}
	//Initialize Router
	router.Initialize()

}
