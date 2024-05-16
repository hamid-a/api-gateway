package main

import (
	"fmt"

	"github.com/hamid-a/api-gateway/internal/config"
	_ "go.uber.org/automaxprocs"
)

func main() {
	fmt.Println("Hello from Api-Gateway. Application is starting...")
	_, err := config.NewConfig("configs/configs.yaml")
	if err != nil {
		panic(err)
	}
}
