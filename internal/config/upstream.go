package config

import "time"

type Upstream struct {
	Name     string    `koanf:"name"`
	Backends []Backend `koanf:"backends"`
}

type GrpcConn struct{}
type HttpConn struct{
	Timeout time.Duration `koanf:"timeout"`
}

type Backend struct {
	Connection string `koanf:"connection"`
	Addr       string `koanf:"addr"`
	GrpcConn
	HttpConn
}
