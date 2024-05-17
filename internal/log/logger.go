package log

import (
	"encoding/json"

	"github.com/hamid-a/api-gateway/internal/config"
	"go.uber.org/zap"
)

func NewLogger(config config.Config) *zap.Logger {
	logLevel := "warn"
	if config.App.Debug {
		logLevel = config.App.LogLevel
	}

	rawJSON := []byte(`{
		"level": "` + logLevel + `",
		"encoding": "json",
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase",
		  "timeKey": "time",
		  "timeEncoder": "RFC3339"
		}
	  }`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	l := zap.Must(cfg.Build())
	defer l.Sync()

	return l
}
