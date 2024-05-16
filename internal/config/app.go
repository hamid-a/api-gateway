package config

type App struct {
	Debug           bool `koanf:"true"`
	ApllicationPort int  `koanf:"applicationPort"`
	MetricsPort     int  `koanf:"metricsPort"`
}
