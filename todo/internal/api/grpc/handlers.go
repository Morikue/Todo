package grpc

import (
	"context"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/protobuf/types/known/emptypb"
	"todo/internal/models"
	"todo/pkg/ctxutil"
	todo "todo/pkg/grpc_stubs/todos"
)

func (s *server) CreateToDo(ctx context.Context, todo *todo.ShortTodoDTO) (*todo.FullTodoDTO, error) {
	ctx = ctxutil.SetRequestIdFromGrpcToContext(ctx)

	span, ctx := opentracing.StartSpanFromContext(ctx, "grpc_handler.CreateToDo")
	defer span.Finish()

	requestId, _ := ctxutil.GetRequestIDFromContext(ctx)
	s.logger.Info().
		Str("requestId", requestId).
		Msgf("request id %s", requestId)

	newTodo, err := models.NewEmptyTodoDTO().FromGRPCShortWithNewId(todo)
	if err != nil {
		return nil, err
	}

	response, err := s.todoService.CreateToDo(ctx, newTodo)
	if err != nil {
		return nil, err
	}

	return response.ToGRPCFull(), nil
}

func (s *server) UpdateToDo(ctx context.Context, todo *todo.ShortTodoDTO) (*todo.FullTodoDTO, error) {
	ctx = ctxutil.SetRequestIdFromGrpcToContext(ctx)

	span, ctx := opentracing.StartSpanFromContext(ctx, "grpc_handler.UpdateToDo")
	defer span.Finish()

	newTodo, err := models.NewEmptyTodoDTO().FromGRPCShort(todo)
	if err != nil {
		return nil, err
	}

	response, err := s.todoService.UpdateToDo(ctx, newTodo)
	if err != nil {
		return nil, err
	}

	return response.ToGRPCFull(), nil
}

func (s *server) GetTodoById(ctx context.Context, todoId *todo.TodoID) (*todo.FullTodoDTO, error) {
	ctx = ctxutil.SetRequestIdFromGrpcToContext(ctx)

	span, ctx := opentracing.StartSpanFromContext(ctx, "grpc_handler.GetTodoById")
	defer span.Finish()

	id, err := uuid.Parse(todoId.Id)
	if err != nil {
		return nil, err
	}

	response, err := s.todoService.GetToDo(ctx, id)
	if err != nil {
		return nil, err
	}

	return response.ToGRPCFull(), nil
}

func (s *server) GetToDos(ctx context.Context, todosRequest *todo.GetTodosRequest) (*todo.GetTodosResponse, error) {
	ctx = ctxutil.SetRequestIdFromGrpcToContext(ctx)

	span, ctx := opentracing.StartSpanFromContext(ctx, "grpc_handler.GetToDos")
	defer span.Finish()

	request := models.NewEmptyGetTodosDTO().FromGRPCRequest(todosRequest)
	response, err := s.todoService.GetToDos(ctx, request)
	if err != nil {
		return nil, err
	}

	return models.SliceToGRPCResponse(response), nil
}

func (s *server) DeleteTodo(ctx context.Context, todoId *todo.TodoID) (*emptypb.Empty, error) {
	ctx = ctxutil.SetRequestIdFromGrpcToContext(ctx)

	span, ctx := opentracing.StartSpanFromContext(ctx, "grpc_handler.DeleteTodo")
	defer span.Finish()

	id, err := uuid.Parse(todoId.Id)
	if err != nil {
		return nil, err
	}

	err = s.todoService.DeleteToDo(ctx, id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
