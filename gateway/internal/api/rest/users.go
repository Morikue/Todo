package rest

import (
	"encoding/json"
	"errors"
	appErrors "gateway/internal/app_errors"
	"gateway/internal/models"
	"gateway/pkg/ctxutil"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"net/http"
	"strconv"
)

func (h *GatewayHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId, _ := ctxutil.GetRequestIDFromContext(ctx)
	span, ctx := opentracing.StartSpanFromContext(ctx, "gateway.RegisterUser")
	defer span.Finish()

	// Обработка запроса на регистрацию нового пользователя.
	var newUser = new(models.CreateUserDTO)
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		h.logger.Error().
			Str("requestId", requestId).
			Msgf("[RegisterUser] unmarshal:%s", err)
		h.ErrorBadRequest(w)
		return
	}

	// передаем данные в слой сервиса
	userID, err := h.gatewayService.RegisterUser(ctx, newUser)
	if err != nil {
		if errors.As(err, &appErrors.ErrUsernameOrEmailIsUsed) {
			h.ErrorUsernameOrEmailAlreadyUsed(w)
			return
		}

		h.logger.Error().
			Str("requestId", requestId).
			Msgf("[RegisterUser] register:%s", err)
		h.ErrorInternalApi(w)
		return
	}

	// упаковываем данные для передачи пользователю
	response := struct {
		UserID int `json:"user_id"`
	}{UserID: userID}

	h.JSONSuccessRespond(w, response)
}

func (h *GatewayHandler) GetGetUserById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId, _ := ctxutil.GetRequestIDFromContext(ctx)
	span, ctx := opentracing.StartSpanFromContext(ctx, "gateway.GetGetUserById")
	defer span.Finish()

	// Обработка запроса на получение информации о пользователе.
	userID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		h.logger.Error().
			Str("requestId", requestId).
			Msgf("[GetGetUserById] get id from query:%s", err)
		h.ErrorBadRequest(w)
		return
	}

	// передаем данные в слой сервиса
	user, err := h.gatewayService.GetUserByID(ctx, userID)
	if err != nil {
		if errors.As(err, &appErrors.ErrNotFound) {
			h.ErrorNotFound(w)
			return
		}

		h.logger.Error().
			Str("requestId", requestId).
			Msgf("[GetGetUserById] get user:%s", err)
		h.ErrorInternalApi(w)
		return
	}

	// возвращаем данные пользователю
	h.JSONSuccessRespond(w, user)
}

func (h *GatewayHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId, _ := ctxutil.GetRequestIDFromContext(ctx)
	span, ctx := opentracing.StartSpanFromContext(ctx, "gateway.UpdateUser")
	defer span.Finish()

	// Обработка запроса на обновление информации о пользователе.
	var updatedUser = new(models.UserDTO)
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		h.logger.Error().
			Str("requestId", requestId).
			Msgf("[UpdateUser] unmarshall: %s", err)
		h.ErrorBadRequest(w)
		return
	}

	// передаем данные в слой сервиса
	response, err := h.gatewayService.UpdateUser(ctx, updatedUser)
	if err != nil {
		h.logger.Error().
			Str("requestId", requestId).
			Msgf("[UpdateUser] update user: %s", err)
		h.ErrorInternalApi(w)
		return
	}

	// передаем ответ пользователю
	h.JSONSuccessRespond(w, response)
}

func (h *GatewayHandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId, _ := ctxutil.GetRequestIDFromContext(ctx)
	span, ctx := opentracing.StartSpanFromContext(ctx, "gateway.UpdatePassword")
	defer span.Finish()

	// Обработка запроса на изменение пароля пользователя.
	var passwordRequest = new(models.UpdateUserPasswordDTO)
	if err := json.NewDecoder(r.Body).Decode(&passwordRequest); err != nil {
		h.logger.Error().
			Str("requestId", requestId).
			Msgf("[UpdatePassword] unmarshall: %s", err)
		h.ErrorBadRequest(w)
		return
	}

	// передаем данные в слой сервиса
	if err := h.gatewayService.UpdatePassword(ctx, passwordRequest); err != nil {
		h.logger.Error().
			Str("requestId", requestId).
			Msgf("[UpdatePassword] update password: %s", err)
		h.ErrorInternalApi(w)
		return
	}

	// возвращаем пользователю ответ - в данном случе просто status 200
	h.JSONSuccessRespond(w, nil)
}

func (h *GatewayHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId, _ := ctxutil.GetRequestIDFromContext(ctx)
	span, ctx := opentracing.StartSpanFromContext(ctx, "gateway.DeleteUser")
	defer span.Finish()

	// Обработка запроса на удаление пользователя.
	userID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		h.logger.Error().
			Str("requestId", requestId).
			Msgf("[DeleteUser] get id from query: %s", err)
		h.ErrorBadRequest(w)
		return
	}

	// передаем данные в слой сервиса
	if err := h.gatewayService.DeleteUser(ctx, userID); err != nil {
		h.logger.Error().
			Str("requestId", requestId).
			Msgf("[DeleteUser] delete user: %s", err)
		h.ErrorInternalApi(w)
		return
	}

	// возвращаем пользователю ответ - в данном случе просто status 200
	h.JSONSuccessRespond(w, nil)
}

func (h *GatewayHandler) UserLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId, _ := ctxutil.GetRequestIDFromContext(ctx)
	span, ctx := opentracing.StartSpanFromContext(ctx, "gateway.UserLogin")
	defer span.Finish()

	// Обработка запроса на удаление пользователя.
	var request = new(models.UserLoginDTO)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.logger.Error().
			Str("requestId", requestId).
			Msgf("[UserLogin] unmarshall: %s", err)
		h.ErrorBadRequest(w)
		return
	}

	response, err := h.gatewayService.Login(ctx, request)
	if err != nil {
		if errors.As(err, &appErrors.ErrWrongCredentials) {
			h.ErrorWrongCredentials(w)
			return
		}

		h.logger.Error().
			Str("requestId", requestId).
			Msgf("[UserLogin] login: %s", err)
		h.ErrorInternalApi(w)
		return
	}

	// возвращаем пользователю ответ
	h.JSONSuccessRespond(w, response)
}

func (h *GatewayHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestId, _ := ctxutil.GetRequestIDFromContext(ctx)
	span, ctx := opentracing.StartSpanFromContext(ctx, "gateway.Refresh")
	defer span.Finish()

	// Обработка запроса на удаление пользователя.
	var request = new(models.UserTokens)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.logger.Error().
			Str("requestId", requestId).
			Msgf("[Refresh] unmarshall: %s", err)
		h.ErrorBadRequest(w)
		return
	}

	// передаем данные слою бизнес-логики
	response, err := h.gatewayService.Refresh(ctx, request.RefreshToken)
	if err != nil {
		h.logger.Error().
			Str("requestId", requestId).
			Msgf("[Refresh] refresh: %s", err)
		h.ErrorInternalApi(w)
		return
	}

	// возвращаем пользователю ответ
	h.JSONSuccessRespond(w, response)
}
