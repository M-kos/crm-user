package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	DefaultLocalEnvFile = ".env"
	EnvPrefix           = "crm"
	LocalEnvFileKey     = "CRM_ENV_FILE"
)

type Config struct {
	Port string `envconfig:"PORT"`
}

func New() *Config {
	var envFile string
	envFile = DefaultLocalEnvFile

	value := os.Getenv(LocalEnvFileKey)

	if value != "" {
		envFile = value
	}

	err := godotenv.Load(envFile)
	if err != nil {
		slog.Error(err.Error())
	}

	var c Config

	err = envconfig.Process(EnvPrefix, &c)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	return &c
}
