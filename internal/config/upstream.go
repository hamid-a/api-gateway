package config

type Upstreams struct {
	Name    string    `koanf:"name"`
	Backend []Backend `koanf:"backend"`
}

type GrpcConn struct{}
type HttpConn struct{}

type Backend struct {
	Connection string `koanf:"connection"`
	GrpcConn
	HttpConn
}
