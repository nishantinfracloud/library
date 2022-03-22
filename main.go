package main

import (
	"library/internal/app/server"
	"os"

	"library/pkg/config"
	"library/pkg/constants"
)

func main() {
	service := "library"
	environment := os.Getenv("BOOT_CUR_ENV")
	if environment == "" {
		environment = constants.DevEnvironment
	}

	// initialize config
	config.Init(service, environment, constants.ConfigFilePath)

	// initialize server
	server.Init()
}
