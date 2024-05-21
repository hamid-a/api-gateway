package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) Proxy(c *gin.Context) {
	upstream, exists := handler.upstream[c.GetString("upstream")]

	if !exists {
		handler.logger.Sugar().Error("service unavailable", "upstream", c.GetString("upstream"))
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"error": "service unavailable"})
		return
	}

	err := upstream.Forward(c)

	if err != nil {
		handler.logger.Sugar().Error("service unavailable", "upstream", c.GetString("upstream"), "error", err.Error())
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
	}
}
