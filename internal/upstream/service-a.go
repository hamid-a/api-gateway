package upstream

import (
	"fmt"
	"github.com/gin-gonic/gin"
	config "github.com/hamid-a/api-gateway/internal/config"
	"io"
	"net/http"
	"sync"
)

type Backend struct {
	BaseURL    string
	HTTPClient *http.Client
}

type ServiceA struct {
	Backend []Backend
	index   int
	mu      sync.Mutex
}

func NewServiceA(c config.Upstream) *ServiceA {
	service := ServiceA{}
	for _, v := range c.Backends {
		b := Backend{
			BaseURL:    v.Addr,
			HTTPClient: &http.Client{Timeout: v.Timeout},
		}
		service.Backend = append(service.Backend, b)
	}

	return &service
}

// GetBackend returns a backend of up stream with roundrobin algorithm
func (upstream *ServiceA) getBackend() (*Backend, error) {
	upstream.mu.Lock()
	defer upstream.mu.Unlock()

	if len(upstream.Backend) == 0 {
		return nil, fmt.Errorf("no backends available for upstream")
	}
	// check circute breaker

	backend := upstream.Backend[upstream.index]
	upstream.index = (upstream.index + 1) % len(upstream.Backend)
	return &backend, nil
}

func (upstream *ServiceA) Forward(c *gin.Context) {
	backend, err := upstream.getBackend()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "no available upstream"})
		return
	}

	url := c.GetString("path")
	req, err := http.NewRequest(c.Request.Method, backend.BaseURL+url, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req.Header = c.Request.Header
	resp, err := backend.HTTPClient.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			c.Header(key, value)
		}
	}

	c.Status(resp.StatusCode)
	if _, err := io.Copy(c.Writer, resp.Body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
