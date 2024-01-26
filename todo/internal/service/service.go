package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog"
	"time"
	"todo/config"
	"todo/internal/models"
	"todo/pkg/ctxutil"
)

type TodoService struct {
	todoRepo           TodoRepository
	cfg                *config.Config
	logger             *zerolog.Logger
	todoRabbitProducer RabbitProducer
	userServiceClient  UsersServiceClient
}

func NewTodoService(
	cfg *config.Config,
	todoRepo TodoRepository,
	logger *zerolog.Logger,
	userServiceClient UsersServiceClient,
	todoRabbitProducer RabbitProducer,
) *TodoService {
	return &TodoService{
		todoRepo:           todoRepo,
		cfg:                cfg,
		userServiceClient:  userServiceClient,
		todoRabbitProducer: todoRabbitProducer,
		logger:             logger,
	}
}

func (s *TodoService) CreateToDo(ctx context.Context, newTodo *models.TodoDTO) (*models.TodoDTO, error) {
	ctx = ctxutil.SetRequestIdFromGrpcToContext(ctx)

	span, ctx := opentracing.StartSpanFromContext(ctx, "service.CreateToDo")
	defer span.Finish()

	createdTodo, err := s.todoRepo.CreateToDo(ctx, newTodo.ToDAO())
	if err != nil {
		return nil, fmt.Errorf("[CreateToDo] create todo: %w", err)
	}

	var receivers = make([]string, 0, 2)

	user, err := s.userServiceClient.GetUserByID(ctx, createdTodo.Assignee)
	if err != nil {
		return nil, fmt.Errorf("[CreateToDo] get user by id:%w", err)
	}
	receivers = append(receivers, user.Email)

	if createdTodo.Assignee != createdTodo.CreatedBy {
		userCreatedBy, err := s.userServiceClient.GetUserByID(ctx, createdTodo.CreatedBy)
		if err != nil {
			return nil, fmt.Errorf("[CreateToDo] get user by id:%w", err)
		}
		receivers = append(receivers, userCreatedBy.Email)
	}

	data, err := json.Marshal(models.TodoMailItem{
		TodoEventType: models.TodoEventTypeCreateTodo,
		Receivers:     receivers,
		AssigneeName:  user.Username,
		Description:   createdTodo.Description,
	})
	if err != nil {
		return nil, fmt.Errorf("[CreateToDo] marshal new todo mssg:%w", err)
	}

	requestID, _ := ctxutil.GetRequestIDFromContext(ctx)
	err = s.todoRabbitProducer.Publish(data, requestID)
	if err != nil {
		return nil, fmt.Errorf("[CreateToDo] publish new todo letter mssg:%w", err)
	}

	return createdTodo.ToDTO(), nil
}

func (s *TodoService) UpdateToDo(ctx context.Context, newTodo *models.TodoDTO) (*models.TodoDTO, error) {
	ctx = ctxutil.SetRequestIdFromGrpcToContext(ctx)

	span, ctx := opentracing.StartSpanFromContext(ctx, "service.UpdateToDo")
	defer span.Finish()

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

	user, err := s.userServiceClient.GetUserByID(ctx, existedTodo.Assignee)
	if err != nil {
		return nil, fmt.Errorf("[UpdateToDo] get user by id:%w", err)
	}

	data, err := json.Marshal(models.TodoMailItem{
		TodoEventType: models.TodoEventTypeUpdateTodo,
		Receivers:     []string{user.Email},
		AssigneeName:  user.Username,
		Description:   existedTodo.Description,
	})
	if err != nil {
		return nil, fmt.Errorf("[UpdateToDo] marshal update todo mssg:%w", err)
	}

	requestID, _ := ctxutil.GetRequestIDFromContext(ctx)
	err = s.todoRabbitProducer.Publish(data, requestID)
	if err != nil {
		return nil, fmt.Errorf("[UpdateToDo] publish update todo letter mssg:%w", err)
	}

	return response.ToDTO(), nil
}

func (s *TodoService) GetToDos(ctx context.Context, todos *models.GetTodosDTO) ([]models.TodoDTO, error) {
	ctx = ctxutil.SetRequestIdFromGrpcToContext(ctx)

	span, ctx := opentracing.StartSpanFromContext(ctx, "service.GetToDos")
	defer span.Finish()

	existedTodos, err := s.todoRepo.GetToDos(ctx, todos)
	if err != nil {
		return nil, fmt.Errorf("[GetToDos] get todos: %w", err)
	}

	response := models.SliceDAOToDTO(&existedTodos)

	return *response, nil
}

func (s *TodoService) GetToDo(ctx context.Context, todoID uuid.UUID) (*models.TodoDTO, error) {
	ctx = ctxutil.SetRequestIdFromGrpcToContext(ctx)

	span, ctx := opentracing.StartSpanFromContext(ctx, "service.GetToDo")
	defer span.Finish()

	response, err := s.todoRepo.GetToDo(ctx, todoID)
	if err != nil {
		return nil, fmt.Errorf("[GetToDo] get todo: %w", err)
	}

	return response.ToDTO(), nil
}

func (s *TodoService) DeleteToDo(ctx context.Context, todoID uuid.UUID) error {
	ctx = ctxutil.SetRequestIdFromGrpcToContext(ctx)

	span, ctx := opentracing.StartSpanFromContext(ctx, "service.DeleteToDo")
	defer span.Finish()

	existedTodo, err := s.todoRepo.GetToDo(ctx, todoID)
	if err != nil {
		return fmt.Errorf("[DeleteToDo] get todo: %w", err)
	}

	err = s.todoRepo.DeleteToDo(ctx, todoID)
	if err != nil {
		return fmt.Errorf("[DeleteToDo] delete todo: %w", err)
	}

	user, err := s.userServiceClient.GetUserByID(ctx, existedTodo.Assignee)
	if err != nil {
		return fmt.Errorf("[DeleteToDo] get user by id:%w", err)
	}

	data, err := json.Marshal(models.TodoMailItem{
		TodoEventType: models.TodoEventTypeDeleteTodo,
		Receivers:     []string{user.Email},
		AssigneeName:  user.Username,
		Description:   existedTodo.Description,
	})
	if err != nil {
		return fmt.Errorf("[DeleteToDo] marshal delete todo mssg:%w", err)
	}

	requestID, _ := ctxutil.GetRequestIDFromContext(ctx)
	err = s.todoRabbitProducer.Publish(data, requestID)
	if err != nil {
		return fmt.Errorf("[DeleteToDo] publish delete todo letter mssg:%w", err)
	}

	return nil
}
