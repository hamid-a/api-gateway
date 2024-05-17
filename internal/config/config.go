package config

import (
	"fmt"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	App       App         `koanf:"app"`
	Upstreams []Upstreams `koanf:"upstreams"`
}

// NewConfig returns a new Config
func NewConfig(path string) (Config, error) {
	config, err := loadConfig(path)
	if err != nil {
		return config, fmt.Errorf("error in reading configs, err: %w", err)
	}

	return config, err
}

// LoadConfig loads the application configuration from the specified config file's path and name.
func loadConfig(configPath string) (Config, error) {
	k := koanf.New(".")

	if err := k.Load(file.Provider(configPath), yaml.Parser()); err != nil {
		return Config{}, err
	}

	k.Load(env.Provider("API_GATEWAY_", ".", func(s string) string {
		return strings.ReplaceAll(strings.ToLower(strings.TrimPrefix(s, "API_GATEWAY_")), "_", ".")
	}), nil)

	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
