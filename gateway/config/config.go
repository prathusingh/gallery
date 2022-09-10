package config

import (
	"flag"
	"os"

	"github.com/prathusingh/gallery/pkg/constants"
	"github.com/prathusingh/gallery/pkg/logger"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "API Gateway microservice config path")
}

type Config struct {
	ServiceName string
	Logger      *logger.Config
}

// InitConfig initializes the config
func InitConfig() (*Config, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv(constants.ConfigPath)
		configPath = configPathFromEnv
	}
	cfg := &Config{}

	return cfg, nil
}
