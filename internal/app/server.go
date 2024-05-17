package app

import (
	"github.com/hamid-a/api-gateway/internal/config"
	"go.uber.org/zap"
	"net/http"
	"fmt"
)

func InitServer(config config.Config, logger *zap.Logger) *http.Server {
	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", config.App.AppPort),
	}

	go func() {
		logger.Sugar().Infof("HTTP server is up and running on 0.0.0.0:%d", config.App.AppPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Sugar().Fatal("error on initializing server", "err", err)
		}
	}()

	return srv
}
