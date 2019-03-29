package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

// ServerConfig holds the service configuration
type Config struct {
	LogDebug bool `envconfig:"LOG_DEBUG" default:"false"`
	GrpcPort int  `envconfig:"GRPC_PORT" default:"50051"`
}

// NewConfig returns a new ServerConfig. The configuration is parsed from environment variables.
// Default values are only set if an environment variable is not set
func NewConfig() *Config {
	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		logrus.WithError(err).Fatal("unable to process configuration")
	}

	return &cfg
}
