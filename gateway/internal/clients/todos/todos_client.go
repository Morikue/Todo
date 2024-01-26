package todos

import (
	"context"
	"fmt"
	"gateway/config"
	"gateway/internal/models"
	"gateway/pkg/ctxutil"
	todo "gateway/pkg/grpc_stubs/todos"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type TodosClient struct {
	client todo.TodoServiceClient
}

func NewTodosClient(cfg *config.Config, logger *zerolog.Logger) (*TodosClient, error) {
	appAddr := fmt.Sprintf("%s:%s", cfg.TodosClient.AppHost, cfg.TodosClient.AppGrpcPort)

	logger.Info().Msgf("[NewTodosClient] connecting via GRPC to todos at %s", appAddr)

	conn, err := grpc.Dial(
		appAddr,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("[NewTodosClient] connect to todos GRPC service: %w", err)
	}

	client := todo.NewTodoServiceClient(conn)
	return &TodosClient{
		client: client,
	}, nil
}

func (c *TodosClient) CreateToDo(ctx context.Context, newTodo *models.CreateTodoDTO) (*models.TodoDTO, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "client.CreateToDo")
	defer span.Finish()

	ctx = ctxutil.SetRequestIdFromContextToGrpc(ctx)
	todoItem, err := c.client.CreateToDo(ctx, newTodo.ToGRPCShort())
	if err != nil {
		return nil, fmt.Errorf("[CreateToDo] create todo: %w", err)
	}

	response, err := models.NewEmptyTodoDTO().FromGRPCFull(todoItem)
	if err != nil {
		return nil, fmt.Errorf("[CreateToDo] get dto from grpc: %w", err)
	}

	return response, nil
}

func (c *TodosClient) UpdateToDo(ctx context.Context, newTodo *models.TodoDTO) (*models.TodoDTO, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "client.UpdateToDo")
	defer span.Finish()

	ctx = ctxutil.SetRequestIdFromContextToGrpc(ctx)
	todoItem, err := c.client.UpdateToDo(ctx, newTodo.ToGRPCShort())
	if err != nil {
		return nil, fmt.Errorf("[CreateToDo] updating: %w", err)
	}

	response, err := models.NewEmptyTodoDTO().FromGRPCFull(todoItem)
	if err != nil {
		return nil, fmt.Errorf("[UpdateToDo] get dto from grpc: %w", err)
	}

	return response, nil
}

func (c *TodosClient) GetToDos(ctx context.Context, todos *models.GetTodosDTO) ([]models.TodoDTO, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "client.GetToDos")
	defer span.Finish()

	ctx = ctxutil.SetRequestIdFromContextToGrpc(ctx)
	storedTodos, err := c.client.GetToDos(ctx, todos.ToGRPCRequest())
	if err != nil {
		return nil, fmt.Errorf("[GetToDos] get: %w", err)
	}

	response, err := models.SliceFromGRPCResponse(storedTodos)
	if err != nil {
		return nil, fmt.Errorf("[GetToDos] get dto from grpc: %w", err)
	}

	return response, nil
}

func (c *TodosClient) GetToDo(ctx context.Context, todoID uuid.UUID) (*models.TodoDTO, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "client.GetToDo")
	defer span.Finish()

	ctx = ctxutil.SetRequestIdFromContextToGrpc(ctx)
	todoItem, err := c.client.GetTodoById(ctx, &todo.TodoID{
		Id: todoID.String(),
	})
	if err != nil {
		return nil, fmt.Errorf("[GetToDo] get: %w", err)
	}

	response, err := models.NewEmptyTodoDTO().FromGRPCFull(todoItem)
	if err != nil {
		return nil, fmt.Errorf("[CreateToDo] get dto from grpc: %w", err)
	}

	return response, nil
}

func (c *TodosClient) DeleteToDo(ctx context.Context, todoID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "client.DeleteToDo")
	defer span.Finish()

	ctx = ctxutil.SetRequestIdFromContextToGrpc(ctx)
	_, err := c.client.DeleteTodo(ctx, &todo.TodoID{
		Id: todoID.String(),
	})
	if err != nil {
		return fmt.Errorf("[DeleteToDo] delete: %w", err)
	}

	return nil
}
