package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hamid-a/api-gateway/internal/config"
	"github.com/hamid-a/api-gateway/internal/upstream"
	"go.uber.org/zap"
)

type Handler struct {
	upstream upstream.UpStream
	config   config.Config
	logger   *zap.Logger
}

func NewHandler(c config.Config, u upstream.UpStream, l *zap.Logger) Handler {
	return Handler{upstream: u, config: c, logger: l}
}

func (handler *Handler) InitRoutes(e *gin.Engine) {
	// Define routes
	for _, rule := range handler.config.Rules {
		for _, method := range rule.Methods {
			switch method {
			case "GET":
				e.GET(rule.Path, handler.Proxy)
			case "POST":
				e.POST(rule.Path, handler.Proxy)
			case "OPTIONS":
				e.OPTIONS(rule.Path, handler.Proxy)
			}
		}
	}
}
