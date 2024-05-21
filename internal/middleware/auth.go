package middleware

import (
	"github.com/gin-gonic/gin"

)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetBool("auth") {
			// Impelemt auth logic
		}
		c.Next()
	}
}
