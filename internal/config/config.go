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

type PostgresConfig struct {
	User     string `envconfig:"CRM_USER_POSTGRES_USER"`
	Password string `envconfig:"CRM_USER_POSTGRES_PASSWORD"`
	Name     string `envconfig:"CRM_USER_POSTGRES_NAME"`
	Port     string `envconfig:"CRM_USER_POSTGRES_PORT"`
	Host     string `envconfig:"CRM_USER_POSTGRES_HOST"`
}

type Config struct {
	Port     int            `envconfig:"CRM_USER_PORT"`
	Postgres PostgresConfig `envconfig:"POSTGRES"`
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
