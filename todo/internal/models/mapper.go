package models

import (
	"fmt"
	"github.com/google/uuid"
	ts "google.golang.org/protobuf/types/known/timestamppb"
	"time"
	"todo/pkg/grpc_stubs/todos"
)

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

func (d *TodoDTO) FromGRPCShortWithNewId(dto *todo.ShortTodoDTO) (*TodoDTO, error) {
	newId := uuid.New()

	return &TodoDTO{
		ID:          newId,
		CreatedBy:   int(dto.CreatedBy),
		Assignee:    int(dto.Assignee),
		Description: dto.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
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

func (d *TodoDTO) ToDAO() *TodoDAO {
	return &TodoDAO{
		ID:          d.ID,
		CreatedBy:   d.CreatedBy,
		Assignee:    d.Assignee,
		Description: d.Description,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}

func (d *TodoDAO) ToDTO() *TodoDTO {
	return &TodoDTO{
		ID:          d.ID,
		CreatedBy:   d.CreatedBy,
		Assignee:    d.Assignee,
		Description: d.Description,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}

func SliceDAOToDTO(daoSlice *[]TodoDAO) *[]TodoDTO {
	dtos := make([]TodoDTO, len(*daoSlice))
	for i, dto := range *daoSlice {
		dtos[i] = *dto.ToDTO()
	}

	return &dtos
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

func (d *GetTodosDTO) FromGRPCRequest(req *todo.GetTodosRequest) *GetTodosDTO {
	return &GetTodosDTO{
		CreatedBy: int(req.CreatedBy),
		Assignee:  int(req.Assignee),
		DateFrom:  req.DateFrom.AsTime(),
		DateTo:    req.DateTo.AsTime(),
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

func SliceToGRPCResponse(slice []TodoDTO) *todo.GetTodosResponse {
	var dtoSlice = todo.GetTodosResponse{
		Items: make([]*todo.FullTodoDTO, len(slice)),
	}

	for i, value := range slice {
		var newDto = todo.FullTodoDTO{
			Id:          value.ID.String(),
			CreatedBy:   int32(value.CreatedBy),
			Assignee:    int32(value.Assignee),
			Description: value.Description,
			CreatedAt:   ts.New(value.CreatedAt),
			UpdatedAt:   ts.New(value.UpdatedAt),
		}
		dtoSlice.Items[i] = &newDto
	}

	return &todo.GetTodosResponse{}
}
