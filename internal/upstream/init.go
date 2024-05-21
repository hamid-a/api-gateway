package upstream

import (
	"github.com/gin-gonic/gin"
	"github.com/hamid-a/api-gateway/internal/config"
)

type Service interface {
	Forward(c *gin.Context) error
	getBackend() (*Backend, error)
}

type UpStream map[string]Service

func Init(c config.Config) (UpStream, error) {
	upstreams := make(map[string]Service)
	for _, u := range c.Upstreams {
		if u.Name == "ServiceA" {
			upstreams["ServiceA"] = NewServiceA(u)
		} else if u.Name == "ServiceB" {
			upstreams["ServiceB"] = NewServiceB(u)
		}
	}

	return upstreams, nil
}
