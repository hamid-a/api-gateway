package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hamid-a/api-gateway/internal/config"
	"github.com/hamid-a/api-gateway/internal/upstream"
)

type Handler struct {
	Upstream upstream.UpStream
	Config   config.Config
}

func NewHandler(Config config.Config, upstream upstream.UpStream) Handler {
	return Handler{Upstream: upstream, Config: Config}
}

func (handler *Handler) InitRoutes(e *gin.Engine) {
	// Define routes
	for _, rule := range handler.Config.Rules {
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
