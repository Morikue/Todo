package config

import (
	"github.com/kelseyhightower/envconfig"
	"todo/pkg/logging"
	"todo/pkg/postgresql"
	"todo/pkg/rabbitmq"
)

type Config struct {
	App          AppConfig             `envconfig:"APP"`
	Grpc         Grpc                  `envconfig:"GRPC"`
	Logging      logging.LoggerConfig  `envconfig:"LOG"`
	Postgres     postgresql.PostgreSQL `envconfig:"POSTGRES"`
	UsersClient  UsersClient           `envconfig:"USERS"`
	RabbitConfig rabbitmq.RabbitConfig `envconfig:"RABBITMQ"`
	TodoExchange string                `envconfig:"RABBITMQ_TODO_EXCHANGE" default:"todo.exchange"`
	TodoQueue    string                `envconfig:"RABBITMQ_TODO_QUEUE" default:"todo.queue"`
}

type Grpc struct {
	AppHost string `envconfig:"GRPC_HOST" required:"true" default:"0.0.0.0"`
	AppPort string `envconfig:"GRPC_PORT" required:"true" default:"50001"`
}

type AppConfig struct {
	AppHost string `envconfig:"APP_HOST" required:"true" default:"0.0.0.0"`
	AppPort string `envconfig:"APP_PORT" required:"true" default:"3001"`
}

type UsersClient struct {
	AppHost     string `envconfig:"USERS_HOST" required:"true" default:"0.0.0.0"`
	AppRestPort string `envconfig:"USERS_REST_PORT" required:"true" default:"3000"`
	AppGrpcPort string `envconfig:"USERS_GRPC_PORT" required:"true" default:"50000"`
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
