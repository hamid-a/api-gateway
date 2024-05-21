package upstream

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	cb "github.com/hamid-a/api-gateway/internal/cb"
	"github.com/hamid-a/api-gateway/internal/config"
	rpc "github.com/hamid-a/api-gateway/pkg/proto/serviceb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"net/http"
	"sync"
)

type ServiceB struct {
	Backend []Backend
	index   int
	mu      sync.Mutex
}

func NewServiceB(c config.Upstream) *ServiceB {
	service := ServiceB{}
	for _, v := range c.Backends {
		conn, err := grpc.NewClient(
			v.Addr,
			grpc.WithKeepaliveParams(*v.Keepalive),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		)
		if err != nil {
			fmt.Printf("could not initialize grpc connection: %v", err)
			continue
		}
		b := Backend{
			baseURL:      v.Addr,
			serviceBStub: rpc.NewServiceBClient(conn),
			cb:           *cb.NewCircuitBreaker(v.Name, v.Cb),
			timeout:      v.Timeout,
		}
		service.Backend = append(service.Backend, b)
	}

	return &service
}

// GetBackend returns a backend of up stream with roundrobin algorithm
func (upstream *ServiceB) getBackend() (*Backend, error) {
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
		if !backend.cb.IsOpen() {
			selectedBackend = backend
			break
		}
	}

	return &selectedBackend, nil
}

func (upstream *ServiceB) Forward(c *gin.Context) error {
	var resErr error

	backend, err := upstream.getBackend()
	if err != nil {
		return errors.New("no available upstream")
	}

	done, err := backend.cb.Allow()
	if err != nil {
		return errors.New("no available upstream")
	}

	defer func() {
		done(resErr == nil)
	}()

	headerReq := make(map[string]*rpc.HeaderValueList)
	for key, values := range c.Request.Header {
		list := make([]string, 0, len(values))
		for _, value := range values {
			list = append(list, value)
		}
		headerReq[key] = &rpc.HeaderValueList{
			List: list,
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), backend.timeout)
	defer cancel()
	req, _ := io.ReadAll(c.Request.Body)
	resp, err := backend.serviceBStub.Server(
		ctx,
		&rpc.Request{
			Header: headerReq,
			Body:   string(req),
		},
		grpc.WaitForReady(true),
	)

	if err != nil {
		resErr = err
		return resErr
	}

	for key, values := range resp.Header {
		for _, value := range values.List {
			c.Header(key, value)
		}
	}

	c.Status(http.StatusOK)
	c.Writer.WriteString(resp.Body)
	return nil
}
