package config

import (
	"github.com/kelseyhightower/envconfig"
	"notifications/pkg/logging"
	rabbitConfig "notifications/pkg/rabbitmq"
	"notifications/pkg/smtp_client"
)

type Config struct {
	Logging       logging.LoggerConfig      `envconfig:"LOG"`
	SmtpConfig    smtp_client.SmtpConfig    `envconfig:"SMTP"`
	RabbitConfig  rabbitConfig.RabbitConfig `envconfig:"RABBITMQ"`
	UsersExchange string                    `envconfig:"RABBITMQ_USERS_EXCHANGE" default:"users.exchange"`
	UsersQueue    string                    `envconfig:"RABBITMQ_USERS_QUEUE" default:"users.queue"`
	TodoExchange  string                    `envconfig:"RABBITMQ_TODO_EXCHANGE" default:"todo.exchange"`
	TodoQueue     string                    `envconfig:"RABBITMQ_TODO_QUEUE" default:"todo.queue"`
}

func NewFromEnv() *Config {
	c := Config{}
	envconfig.MustProcess("", &c)
	return &c
}
