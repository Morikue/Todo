package rest

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"net/http"
	"todo/internal/api"
	"todo/internal/models"
)

type TodoHandler struct {
	logger      *zerolog.Logger
	todoService api.TodoService
}

func NewTodoHandler(
	logger *zerolog.Logger,
	todoService api.TodoService,
) *TodoHandler {
	return &TodoHandler{
		logger:      logger,
		todoService: todoService,
	}
}

func (h *TodoHandler) CreateToDo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var newTodo = new(models.TodoDTO)
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		h.logger.Error().Msgf("[CreateToDo] unmarshal:%s", err)
		h.ErrorBadRequest(w, "Bad request")
		return
	}

	response, err := h.todoService.CreateToDo(ctx, newTodo)
	if err != nil {
		h.logger.Error().Msgf("[CreateToDo] creating:%s", err)
		h.ErrorInternalError(w, "Can't create Todo")
		return
	}

	h.WriteResponse(w, response)
}

func (h *TodoHandler) GetToDo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	todoId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		h.ErrorBadRequest(w, "Id is not valid UUID")
		return
	}

	response, err := h.todoService.GetToDo(ctx, todoId)
	if err != nil {
		h.logger.Error().Msgf("[GetToDo] getting:%s", err)
		h.ErrorInternalError(w, "Can't get Todo")
		return
	}

	h.WriteResponse(w, response)
}

func (h *TodoHandler) GetToDos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var newTodos = new(models.GetTodosDTO)
	if err := json.NewDecoder(r.Body).Decode(&newTodos); err != nil {
		h.logger.Error().Msgf("[GetToDos] unmarshal:%s", err)
		h.ErrorBadRequest(w, "Bad request")
		return
	}

	response, err := h.todoService.GetToDos(ctx, newTodos)
	if err != nil {
		h.logger.Error().Msgf("[GetToDos] getting:%s", err)
		h.ErrorInternalError(w, "Can't get Todos")
		return
	}

	h.WriteResponse(w, response)
}

func (h *TodoHandler) UpdateToDo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var newTodo = new(models.TodoDTO)
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		h.logger.Error().Msgf("[UpdateToDo] unmarshal:%s", err)
		h.ErrorBadRequest(w, "Bad request")
		return
	}

	response, err := h.todoService.UpdateToDo(ctx, newTodo)
	if err != nil {
		h.logger.Error().Msgf("[UpdateToDo] updating:%s", err)
		h.ErrorInternalError(w, "Can't update Todo")
		return
	}

	h.WriteResponse(w, response)
}

func (h *TodoHandler) DeleteToDo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	todoId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		h.ErrorBadRequest(w, "Id is not valid UUID")
		return
	}

	err = h.todoService.DeleteToDo(ctx, todoId)
	if err != nil {
		h.logger.Error().Msgf("[DeleteToDo] deleting:%s", err)
		h.ErrorInternalError(w, "Can't delete Todo")
		return
	}

	h.WriteResponse(w, nil)
}
