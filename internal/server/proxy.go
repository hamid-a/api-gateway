package server

import (
	"github.com/gin-gonic/gin"
)

func (handler *Handler) Proxy(c *gin.Context) {
	upstream := handler.Upstream[c.GetString("upstream")]

	upstream.Forward(c)
}