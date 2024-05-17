package main

import (
	"fmt"

	"github.com/hamid-a/api-gateway/internal/config"
	"github.com/hamid-a/api-gateway/internal/log"
	_ "go.uber.org/automaxprocs"
)

func main() {
	fmt.Println("Hello from Api-Gateway. Application is starting...")
	config, err := config.NewConfig("configs/configs.yaml")
	if err != nil {
		panic(err)
	}

	logger := log.NewLogger(config)
	logger.Info("logger initialized")
}
