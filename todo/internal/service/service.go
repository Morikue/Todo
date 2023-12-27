package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
	"todo/config"
	"todo/internal/models"
)

type TodoService struct {
	todoRepo TodoRepository
	cfg      *config.Config
}

func NewTodoService(cfg *config.Config, todoRepo TodoRepository) *TodoService {
	return &TodoService{
		todoRepo: todoRepo,
		cfg:      cfg,
	}
}

func (s *TodoService) CreateToDo(ctx context.Context, newTodo *models.TodoDTO) (*models.TodoDTO, error) {
	createdTodo, err := s.todoRepo.CreateToDo(ctx, newTodo.ToDAO())
	if err != nil {
		return nil, fmt.Errorf("[CreateToDo] create todo: %w", err)
	}

	return createdTodo.ToDTO(), nil
}

func (s *TodoService) UpdateToDo(ctx context.Context, newTodo *models.TodoDTO) (*models.TodoDTO, error) {
	existedTodo, err := s.todoRepo.GetToDo(ctx, newTodo.ID)
	if err != nil {
		return nil, fmt.Errorf("[UpdateToDo] get todo: %w", err)
	}

	existedTodo.Description = newTodo.Description
	existedTodo.Assignee = newTodo.Assignee
	existedTodo.UpdatedAt = time.Now()

	response, err := s.todoRepo.UpdateToDo(ctx, existedTodo)
	if err != nil {
		return nil, fmt.Errorf("[UpdateToDo] update todo: %w", err)
	}

	return response.ToDTO(), nil
}

func (s *TodoService) GetToDos(ctx context.Context, todos *models.GetTodosDTO) ([]models.TodoDTO, error) {
	existedTodos, err := s.todoRepo.GetToDos(ctx, todos)
	if err != nil {
		return nil, fmt.Errorf("[GetToDos] get todos: %w", err)
	}

	response := models.SliceDAOToDTO(&existedTodos)

	return *response, nil
}

func (s *TodoService) GetToDo(ctx context.Context, todoID uuid.UUID) (*models.TodoDTO, error) {
	response, err := s.todoRepo.GetToDo(ctx, todoID)
	if err != nil {
		return nil, fmt.Errorf("[GetToDo] get todo: %w", err)
	}

	return response.ToDTO(), nil
}

func (s *TodoService) DeleteToDo(ctx context.Context, todoID uuid.UUID) error {
	err := s.todoRepo.DeleteToDo(ctx, todoID)
	if err != nil {
		return fmt.Errorf("[DeleteToDo] delete todo: %w", err)
	}

	return nil
}
