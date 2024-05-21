package app

import (
	"github.com/hamid-a/api-gateway/internal/config"
	"github.com/hamid-a/api-gateway/internal/upstream"
	"go.uber.org/zap"
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hamid-a/api-gateway/internal/middleware"
)

func InitServer(config config.Config, logger *zap.Logger, upstream upstream.UpStream) *http.Server {
	if config.App.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.Rule(config, upstream))

	registerRoutes(engine, config.Rules)

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

func registerRoutes(e *gin.Engine, rules []config.Rule) {
	// Define routes
	for _, rule := range rules {
		for _, method := range rule.Methods {
			switch method {
			case "GET":
				e.GET(rule.Path, func(c *gin.Context) {})
			case "POST":
				e.POST(rule.Path, func(c *gin.Context) {})
			case "OPTIONS":
				e.OPTIONS(rule.Path, func(c *gin.Context) {})
			}
		}
	}
}
