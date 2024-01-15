package api

import (
	"notifications/internal/models"
)

type UserService interface {
	SendUserMessage(item *models.UserMailItem) error
}

type TodoService interface {
	SendTodoMessage(item *models.TodoMailItem) error
}
