package models

import (
	"fmt"
	todo "gateway/pkg/grpc_stubs/todos"
	"github.com/google/uuid"
	ts "google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type CreateTodoDTO struct {
	CreatedBy   int    `json:"created_by" example:"1"`
	Assignee    int    `json:"assignee" example:"2"`
	Description string `json:"description" example:"todo description"`
}

func NewEmptyCreateTodoDTOO() *CreateTodoDTO {
	return &CreateTodoDTO{}
}

func (d *CreateTodoDTO) ToGRPCShort() *todo.ShortTodoDTO {
	return &todo.ShortTodoDTO{
		Id:          "",
		CreatedBy:   int32(d.CreatedBy),
		Assignee:    int32(d.Assignee),
		Description: d.Description,
	}
}

type TodoDTO struct {
	ID          uuid.UUID `json:"id,omitempty" example:"c0e708fa-a7df-4d9f-a1b8-a3bfe63c433c"`
	CreatedBy   int       `json:"created_by" example:"1"`
	Assignee    int       `json:"assignee" example:"2"`
	Description string    `json:"description" example:"todo description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewEmptyTodoDTO() *TodoDTO {
	return &TodoDTO{}
}

func (d *TodoDTO) ToGRPCShort() *todo.ShortTodoDTO {
	return &todo.ShortTodoDTO{
		Id:          d.ID.String(),
		CreatedBy:   int32(d.CreatedBy),
		Assignee:    int32(d.Assignee),
		Description: d.Description,
	}
}

func (d *TodoDTO) ToGRPCFull() *todo.FullTodoDTO {
	return &todo.FullTodoDTO{
		Id:          d.ID.String(),
		CreatedBy:   int32(d.CreatedBy),
		Assignee:    int32(d.Assignee),
		Description: d.Description,
		CreatedAt:   ts.New(d.CreatedAt),
		UpdatedAt:   ts.New(d.UpdatedAt),
	}
}

func (d *TodoDTO) FromGRPCShort(dto *todo.ShortTodoDTO) (*TodoDTO, error) {
	id, err := uuid.Parse(dto.Id)
	if err != nil {
		return nil, fmt.Errorf("[FromGRPCShort] wrong uuid: %w", err)
	}

	return &TodoDTO{
		ID:          id,
		CreatedBy:   int(dto.CreatedBy),
		Assignee:    int(dto.Assignee),
		Description: dto.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (d *TodoDTO) FromGRPCFull(dto *todo.FullTodoDTO) (*TodoDTO, error) {
	id, err := uuid.Parse(dto.Id)
	if err != nil {
		return nil, fmt.Errorf("[FromGRPCShort] wrong uuid: %w", err)
	}

	return &TodoDTO{
		ID:          id,
		CreatedBy:   int(dto.CreatedBy),
		Assignee:    int(dto.Assignee),
		Description: dto.Description,
		CreatedAt:   dto.CreatedAt.AsTime(),
		UpdatedAt:   dto.UpdatedAt.AsTime(),
	}, nil
}

type GetTodosDTO struct {
	CreatedBy int       `json:"created_by" example:"1"`
	Assignee  int       `json:"assignee" example:"2"`
	DateFrom  time.Time `json:"date_from"`
	DateTo    time.Time `json:"date_to"`
}

func NewEmptyGetTodosDTO() *GetTodosDTO {
	return &GetTodosDTO{}
}

func (d *GetTodosDTO) ToGRPCRequest() *todo.GetTodosRequest {
	return &todo.GetTodosRequest{
		CreatedBy: int32(d.CreatedBy),
		Assignee:  int32(d.Assignee),
		DateFrom:  ts.New(d.DateFrom),
		DateTo:    ts.New(d.DateTo),
	}
}

func SliceFromGRPCResponse(response *todo.GetTodosResponse) ([]TodoDTO, error) {
	var dtoSlice = make([]TodoDTO, len(response.Items))

	for i := 0; i < len(response.Items); i++ {
		id, err := uuid.Parse(response.Items[i].Id)
		if err != nil {
			return nil, fmt.Errorf("[SliceFromGRPCResponse] wrong uuid: %w", err)
		}

		var newDto = TodoDTO{
			ID:          id,
			CreatedBy:   int(response.Items[i].CreatedBy),
			Assignee:    int(response.Items[i].Assignee),
			Description: response.Items[i].Description,
			CreatedAt:   response.Items[i].CreatedAt.AsTime(),
			UpdatedAt:   response.Items[i].UpdatedAt.AsTime(),
		}
		dtoSlice[i] = newDto
	}

	return dtoSlice, nil
}
