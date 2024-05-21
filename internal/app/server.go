package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hamid-a/api-gateway/internal/config"
	"github.com/hamid-a/api-gateway/internal/middleware"
	"github.com/hamid-a/api-gateway/internal/server"
	"github.com/hamid-a/api-gateway/internal/upstream"
	"go.uber.org/zap"
)

func InitServer(config config.Config, logger *zap.Logger, upstream upstream.UpStream) *http.Server {
	if config.App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.Rule(config))
	engine.Use(middleware.Auth())

	handler := server.NewHandler(config, upstream)
	handler.InitRoutes(engine)

	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", config.App.AppPort),
		Handler: engine,
	}

	go func() {
		logger.Sugar().Infof("HTTP server is up and running on 0.0.0.0:%d", config.App.AppPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Sugar().Fatal("error on initializing server", "err", err)
		}
	}()

	return srv
}
