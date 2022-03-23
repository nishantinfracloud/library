package main

import (
	"library/internal/app/config"
	"library/internal/app/constants"
	"library/internal/app/database"
	"library/internal/app/server"
	"os"
)

func main() {
	service := "library"
	environment := os.Getenv("BOOT_CUR_ENV")
	if environment == "" {
		environment = constants.DevEnvironment
	}

	// initialize config
	config.Init(service, environment, constants.ConfigFilePath)

	// initialize database
	database.Init()

	// initialize server
	server.Init()
}
