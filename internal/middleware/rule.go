package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hamid-a/api-gateway/internal/config"
	"github.com/hamid-a/api-gateway/internal/util"
	"net/http"
)

func Rule(config config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method

		for _, rule := range config.Rules {
			if path == rule.Path && util.Contains(rule.Methods, method) {
				if rule.Auth {
					Auth()(c)
				}
				// Implement rule logic and load balancing

				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "route not found"})
		c.Abort()
	}
}
