package service

import (
	"context"
	"fmt"
	"gateway/internal/models"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
)

func (s *GatewayService) CreateToDo(ctx context.Context, newTodo *models.CreateTodoDTO) (*models.TodoDTO, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.CreateToDo")
	defer span.Finish()

	todo, err := s.todoServiceClient.CreateToDo(ctx, newTodo)
	if err != nil {
		return nil, fmt.Errorf("[CreateToDo] create todo:%w", err)
	}

	// возвращаем данные в слой хэндлера
	return todo, nil
}

func (s *GatewayService) UpdateToDo(ctx context.Context, newTodo *models.TodoDTO) (*models.TodoDTO, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.UpdateToDo")
	defer span.Finish()

	todo, err := s.todoServiceClient.UpdateToDo(ctx, newTodo)
	if err != nil {
		return nil, fmt.Errorf("[UpdateToDo] update todo:%w", err)
	}

	// возвращаем данные в слой хэндлера
	return todo, nil
}

func (s *GatewayService) GetToDos(ctx context.Context, todos *models.GetTodosDTO) ([]models.TodoDTO, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.GetToDos")
	defer span.Finish()

	storedTodos, err := s.todoServiceClient.GetToDos(ctx, todos)
	if err != nil {
		return nil, fmt.Errorf("[GetToDos] get todos:%w", err)
	}

	// возвращаем данные в слой хэндлера
	return storedTodos, nil
}

func (s *GatewayService) GetToDo(ctx context.Context, todoID uuid.UUID) (*models.TodoDTO, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.GetToDo")
	defer span.Finish()

	todo, err := s.todoServiceClient.GetToDo(ctx, todoID)
	if err != nil {
		return nil, fmt.Errorf("[GetToDo] get todo:%w", err)
	}

	// возвращаем данные в слой хэндлера
	return todo, nil
}

func (s *GatewayService) DeleteToDo(ctx context.Context, todoID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.DeleteToDo")
	defer span.Finish()

	err := s.todoServiceClient.DeleteToDo(ctx, todoID)
	if err != nil {
		return fmt.Errorf("[DeleteToDo] delete todo:%w", err)
	}

	// возвращаем данные в слой хэндлера
	return nil
}
