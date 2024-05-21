package config

type Rule struct {
	Name     string   `koanf:"name"`
	Path     string   `koanf:"path"`
	Auth     bool     `koanf:"auth"`
	Upstream string   `koanf:"upstream"`
	URL      string   `koanf:"url"`
	Methods  []string `koanf:"methods"`
}
