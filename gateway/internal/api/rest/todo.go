package rest

import (
	"encoding/json"
	"errors"
	"gateway/internal/app_errors"
	"gateway/internal/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *GatewayHandler) CreateToDoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var newTodo = new(models.CreateTodoDTO)
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		h.logger.Error().Msgf("[CreateToDoHandler] unmarshal: %s", err)
		h.ErrorBadRequest(w)
		return
	}

	createdTodo, err := h.gatewayService.CreateToDo(ctx, newTodo)
	if err != nil {
		h.logger.Error().Msgf("[CreateToDoHandler] create todo: %s", err)
		h.ErrorInternalApi(w)
		return
	}

	h.JSONSuccessRespond(w, createdTodo)
}

func (h *GatewayHandler) GetToDoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		h.logger.Error().Msgf("[GetToDoHandler] parse id from url: %s", err)
		h.ErrorBadRequest(w)
		return
	}

	todo, err := h.gatewayService.GetToDo(ctx, id)
	if err != nil {
		if errors.As(err, &app_errors.ErrNotFound) {
			h.ErrorNotFound(w)
			return
		}

		h.logger.Error().Msgf("[GetToDoHandler] get todo: %s", err)
		h.ErrorInternalApi(w)
		return
	}

	h.JSONSuccessRespond(w, todo)
}

func (h *GatewayHandler) GetToDosHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var newTodos = new(models.GetTodosDTO)
	if err := json.NewDecoder(r.Body).Decode(&newTodos); err != nil {
		h.logger.Error().Msgf("[GetToDos] unmarshal:%s", err)
		h.ErrorBadRequest(w)
		return
	}

	response, err := h.gatewayService.GetToDos(ctx, newTodos)
	if err != nil {
		h.logger.Error().Msgf("[GetToDos] getting:%s", err)
		h.ErrorInternalApi(w)
		return
	}

	h.JSONSuccessRespond(w, response)

}

func (h *GatewayHandler) UpdateToDoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		h.logger.Error().Msgf("[UpdateToDoHandler] parse id from url: %s", err)
		h.ErrorBadRequest(w)
		return
	}

	todo := new(models.TodoDTO)
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		h.ErrorBadRequest(w)
		return
	}
	todo.ID = id

	updatedTodo, err := h.gatewayService.UpdateToDo(ctx, todo)
	if err != nil {
		if errors.As(err, &app_errors.ErrNotFound) {
			h.ErrorNotFound(w)
			return
		}

		h.logger.Error().Msgf("[UpdateToDoHandler] update todo: %s", err)
		h.ErrorInternalApi(w)
		return
	}

	h.JSONSuccessRespond(w, updatedTodo)
}

func (h *GatewayHandler) DeleteToDoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		h.logger.Error().Msgf("[DeleteToDoHandler] parse id from url: %s", err)
		h.ErrorBadRequest(w)
		return
	}

	err = h.gatewayService.DeleteToDo(ctx, id)
	if err != nil {
		if errors.As(err, &app_errors.ErrNotFound) {
			h.ErrorNotFound(w)
			return
		}

		h.logger.Error().Msgf("[DeleteToDoHandler] delete todo: %s", err)
		h.ErrorInternalApi(w)
		return
	}

	h.JSONSuccessRespond(w, nil)
}
