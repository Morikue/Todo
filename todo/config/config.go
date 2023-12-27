package config

import (
	"github.com/kelseyhightower/envconfig"
	"todo/pkg/logging"
	"todo/pkg/postgresql"
)

type Config struct {
	App      AppConfig             `envconfig:"APP"`
	Grpc     Grpc                  `envconfig:"GRPC"`
	Logging  logging.LoggerConfig  `envconfig:"LOG"`
	Postgres postgresql.PostgreSQL `envconfig:"POSTGRES"`
}

type Grpc struct {
	AppHost string `envconfig:"GRPC_HOST" required:"true" default:"0.0.0.0"`
	AppPort string `envconfig:"GRPC_PORT" required:"true" default:"50001"`
}

type AppConfig struct {
	AppHost string `envconfig:"APP_HOST" required:"true" default:"0.0.0.0"`
	AppPort string `envconfig:"APP_PORT" required:"true" default:"80009"`
}

func NewFromEnv() *Config {
	c := Config{}
	envconfig.MustProcess("", &c)
	return &c
}

type MigrationsConfig struct {
	Postgres postgresql.PostgreSQL `envconfig:"POSTGRES"`
}

func NewMigrationsFromEnv() *MigrationsConfig {
	c := MigrationsConfig{}
	envconfig.MustProcess("", &c)
	return &c
}
