package main

import (
	"fmt"

	"context"
	"github.com/hamid-a/api-gateway/internal/app"
	"github.com/hamid-a/api-gateway/internal/config"
	"github.com/hamid-a/api-gateway/internal/log"
	_ "go.uber.org/automaxprocs"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Hello from Api-Gateway. Application is starting...")
	config, err := config.NewConfig("configs/configs.yaml")
	if err != nil {
		panic(err)
	}

	logger := log.NewLogger(config)
	logger.Info("logger initialized")

	server := app.InitServer(config, logger)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), config.App.GracefullyShutdownTimeout)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Sugar().Fatal("Failed to shutdown:", "error", err)
	}
	logger.Info("server graceful shut down successful")
}
