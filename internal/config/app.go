package config

import "time"

type App struct {
	AppPort                   int           `koanf:"appPort"`
	MetricsPort               int           `koanf:"metricsPort"`
	Debug                     bool          `koanf:"debug"`
	LogLevel                  string        `koanf:"logLevel"`
	GracefullyShutdownTimeout time.Duration `koanf:"gracefullyShutdownTimeout"`
}
