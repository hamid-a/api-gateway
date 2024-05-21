package upstream

import (
	"github.com/hamid-a/api-gateway/internal/config"
	"github.com/gin-gonic/gin"
)

type Service interface {
	Forward(c *gin.Context, url string)
}

type UpStream map[string]Service

func Init(c config.Config) (error, UpStream) {
	upstreams := make(map[string]Service)
	for _, u := range c.Upstreams {
		if u.Name == "ServiceA" {
			upstreams["ServiceA"] = NewServiceA(u)
		} else if  u.Name == "ServiceA" {

		}
	}

	return nil, upstreams
}
