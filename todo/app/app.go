package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
	"todo/config"
	"todo/internal/api"
	"todo/internal/api/grpc"
	"todo/internal/api/rest"
	"todo/internal/repository"
	"todo/internal/service"
	"todo/pkg/logging"
	"todo/pkg/postgresql"
)

type App struct {
	cfg         *config.Config
	logger      *zerolog.Logger
	router      *mux.Router
	todoService api.TodoService
}

func NewApp(cfg *config.Config) (*App, error) {
	// подключимся к базе данных
	databaseConn, err := postgresql.NewPgxConn(&cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}

	todoRepo := repository.NewTodoRepository(databaseConn)

	logger := logging.NewLogger(cfg.Logging)
	todoService := service.NewTodoService(cfg, todoRepo)

	return &App{
		cfg:         cfg,
		logger:      logger,
		todoService: todoService,
	}, nil
}

func (a *App) RunAPI() error {
	group := new(errgroup.Group)
	group.Go(func() error {
		err := rest.NewRestApi(a.cfg, a.logger, a.todoService)
		return fmt.Errorf("[RunApp] run REST: %w", err)
	})

	group.Go(func() error {
		err := grpc.NewGrpcApi(a.cfg, a.logger, a.todoService)
		return fmt.Errorf("[RunApp] run GRPC: %w", err)
	})

	if err := group.Wait(); err != nil {
		return fmt.Errorf("[RunApp] run: %w", err)
	}

	return nil

}
