package config

import (
	"gateway/pkg/jaeger"
	"gateway/pkg/jwtutil"
	"gateway/pkg/logging"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App         App                  `envconfig:"APP"`
	JWT         jwtutil.JWTUtil      `envconfig:"JWT"`
	Logging     logging.LoggerConfig `envconfig:"LOG"`
	Jaeger      jaeger.JaegerConfig  `envconfig:"JAEGER"`
	UsersClient UsersClient          `envconfig:"USERS"`
	TodosClient TodosClient          `envconfig:"TODOS"`
}

type App struct {
	AppHost string `envconfig:"APP_HOST" required:"true" default:"0.0.0.0"`
	AppPort string `envconfig:"APP_PORT" required:"true" default:"3009"`
}

type UsersClient struct {
	AppHost     string `envconfig:"USERS_HOST" required:"true" default:"0.0.0.0"`
	AppRestPort string `envconfig:"USERS_REST_PORT" required:"true" default:"3000"`
	AppGrpcPort string `envconfig:"USERS_GRPC_PORT" required:"true" default:"50000"`
}

type TodosClient struct {
	AppHost     string `envconfig:"TODOS_HOST" required:"true" default:"0.0.0.0"`
	AppRestPort string `envconfig:"TODOS_REST_PORT" required:"true" default:"3001"`
	AppGrpcPort string `envconfig:"TODOS_GRPC_PORT" required:"true" default:"50001"`
}

func NewFromEnv() *Config {
	c := Config{}
	envconfig.MustProcess("", &c)
	return &c
}
