package app_errors

import "errors"

var (
	ErrIncorrectUserEventType = errors.New("incorrect user message type")
	ErrIncorrectTodoEventType = errors.New("incorrect todo message type")
)
