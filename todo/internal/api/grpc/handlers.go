package grpc

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"todo/internal/models"
	todo "todo/pkg/grpc_stubs/todos"
)

func (s *server) CreateToDo(ctx context.Context, todo *todo.ShortTodoDTO) (*todo.FullTodoDTO, error) {
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
	request := models.NewEmptyGetTodosDTO().FromGRPCRequest(todosRequest)
	response, err := s.todoService.GetToDos(ctx, request)
	if err != nil {
		return nil, err
	}

	return models.SliceToGRPCResponse(response), nil
}

func (s *server) DeleteTodo(ctx context.Context, todoId *todo.TodoID) (*emptypb.Empty, error) {
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
