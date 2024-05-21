package upstream

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	cb "github.com/hamid-a/api-gateway/internal/cb"
	config "github.com/hamid-a/api-gateway/internal/config"
)

type Backend struct {
	BaseURL    string
	HTTPClient *http.Client
	CB         cb.CircuitBreaker
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
			CB:         *cb.NewCircuitBreaker(v.Name, v.Cb),
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

	var selectedBackend Backend
	for i := 0; i < len(upstream.Backend); i++ {
		backend := upstream.Backend[upstream.index]
		upstream.index = (upstream.index + 1) % len(upstream.Backend)
		// check circute breaker state and if open get another backend
		if !backend.CB.IsOpen() {
			selectedBackend = backend
			break
		}
	}

	return &selectedBackend, nil
}

func (upstream *ServiceA) Forward(c *gin.Context) error {
	var resErr error

	backend, err := upstream.getBackend()
	if err != nil {
		return errors.New("no available upstream")
	}

	done, err := backend.CB.Allow()
	if err != nil {
		return errors.New("no available upstream")
	}

	defer func() {
		done(resErr == nil)
	}()

	url := c.GetString("path")
	req, err := http.NewRequest(c.Request.Method, backend.BaseURL+url, c.Request.Body)
	if err != nil {
		resErr = errors.New("no available upstream")
		return resErr
	}

	req.Header = c.Request.Header
	resp, err := backend.HTTPClient.Do(req)
	if err != nil {
		resErr = err
		return resErr
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			c.Header(key, value)
		}
	}

	c.Status(resp.StatusCode)
	if _, err := io.Copy(c.Writer, resp.Body); err != nil {
		return err
	}

	return nil
}
