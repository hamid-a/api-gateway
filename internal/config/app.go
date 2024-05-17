package config

type App struct {
	AppPort     int    `koanf:"appPort"`
	MetricsPort int    `koanf:"metricsPort"`
	Debug       bool   `koanf:"debug"`
	LogLevel    string `koanf:"logLevel"`
}
