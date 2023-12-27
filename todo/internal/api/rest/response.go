package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	ErrorMessage string `json:"errorMessage"`
	ErrorCode    int    `json:"errorCode"`
}

func (h *TodoHandler) ErrorBadRequest(w http.ResponseWriter, errorMessage string) {
	h.WriteError(w, ErrorResponse{errorMessage, http.StatusBadRequest})
}

func (h *TodoHandler) ErrorInternalError(w http.ResponseWriter, errorMessage string) {
	h.WriteError(w, ErrorResponse{errorMessage, http.StatusInternalServerError})
}

func (h *TodoHandler) WriteError(w http.ResponseWriter, serverErr ErrorResponse) {
	// установка хэдера ответа
	w.Header().Set("Content-Type", "application/json")

	rawError, err := json.Marshal(serverErr)
	if err != nil {
		// TODO: log
		println("[WriteError] marshal error:", err)
		h.WriteError(w, ErrorResponse{
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: fmt.Sprint("Cannot marshal error response. ", err),
		})
	}

	w.WriteHeader(serverErr.ErrorCode)
	_, err = w.Write(rawError)
	if err != nil {
		// TODO: log
		println("[WriteError] write response error:", err)
	}

}

func (h *TodoHandler) WriteResponse(w http.ResponseWriter, data interface{}) {
	// установка хэдера ответа
	w.Header().Set("Content-Type", "application/json")

	// если нет тела ответа, то просто вернем ответ со статусом 200
	if data == nil {
		w.WriteHeader(http.StatusOK)
		return
	}

	rawData, err := json.Marshal(data)
	if err != nil {
		// TODO: log
		println("[WriteResponse] marshal:", err)
		h.WriteError(w, ErrorResponse{
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: fmt.Sprint("Cannot marshal data. ", err),
		})
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(rawData)
	if err != nil {
		// TODO: log
		println("[WriteResponse] write response:", err)
	}
}
