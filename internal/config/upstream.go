package config

import (
	"google.golang.org/grpc/keepalive"
	"time"
)

type Upstream struct {
	Name     string    `koanf:"name"`
	Backends []Backend `koanf:"backends"`
}

type Backend struct {
	Name       string                      `koanf:"name"`
	Connection string                      `koanf:"connection"`
	Addr       string                      `koanf:"addr"`
	Cb         CircuitBreaker              `koanf:"cb"`
	Keepalive  *keepalive.ClientParameters `koanf:"keepalive"`
	Timeout    time.Duration               `koanf:"timeout"`
}

type CircuitBreaker struct {
	Enabled                bool          `koanf:"enabled"`
	ResetInterval          time.Duration `koanf:"resetInterval"`
	OpenTimeout            time.Duration `koanf:"openTimeout"`
	MaxRequests            uint32        `koanf:"maxRequests"`
	MinRequests            uint32        `koanf:"minRequests"`
	FailureRatioThereshold float64       `koanf:"failureRatioThereshold"`
}
