package rest

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"net/http"
	"todo/config"
	"todo/internal/api"
)

func NewRestApi(
	cfg *config.Config,
	logger *zerolog.Logger,
	todoService api.TodoService,
) error {
	todoHandler := NewTodoHandler(logger, todoService)

	// Настройка роутеров и запуск сервера.
	router := mux.NewRouter()
	router.HandleFunc("/todos", todoHandler.CreateToDo).Methods(http.MethodPost)
	router.HandleFunc("/todos/{id}", todoHandler.GetToDo).Methods(http.MethodGet)
	router.HandleFunc("/todos/batch", todoHandler.GetToDos).Methods(http.MethodPost)
	router.HandleFunc("/todos/{id}", todoHandler.UpdateToDo).Methods(http.MethodPut)
	router.HandleFunc("/todos/{id}", todoHandler.DeleteToDo).Methods(http.MethodDelete)

	appAddr := fmt.Sprintf("%s:%s", cfg.App.AppHost, cfg.App.AppPort) // добавлен
	logger.Info().Msgf("running server at '%s'", appAddr)

	err := http.ListenAndServe(appAddr, router)
	if err != nil {
		return fmt.Errorf("[NewRestApi] listen and serve: %w", err)
	}

	return nil
}
