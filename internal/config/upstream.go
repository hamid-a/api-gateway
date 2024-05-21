package config

import "time"

type Upstream struct {
	Name     string    `koanf:"name"`
	Backends []Backend `koanf:"backends"`
}

type GrpcConn struct{}
type HttpConn struct {
	Timeout time.Duration `koanf:"timeout"`
}

type Backend struct {
	Name       string         `koanf:"name"`
	Connection string         `koanf:"connection"`
	Addr       string         `koanf:"addr"`
	Cb         CircuitBreaker `koanf:"cb"`
	GrpcConn
	HttpConn
}

type CircuitBreaker struct {
	Enabled                bool          `koanf:"enabled"`
	ResetInterval          time.Duration `koanf:"resetInterval"`
	OpenTimeout            time.Duration `koanf:"openTimeout"`
	MaxRequests            uint32        `koanf:"maxRequests"`
	MinRequests            uint32        `koanf:"minRequests"`
	FailureRatioThereshold float64       `koanf:"failureRatioThereshold"`
}
