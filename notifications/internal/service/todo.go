package service

import (
	"fmt"
	"github.com/rs/zerolog"
	"notifications/internal/app_errors"
	"notifications/internal/models"
)

type TodoService struct {
	logger     *zerolog.Logger
	smtpClient SmtpClient
}

func NewTodoService(
	logger *zerolog.Logger,
	smtpClient SmtpClient,
) *TodoService {
	return &TodoService{
		logger:     logger,
		smtpClient: smtpClient,
	}
}

func (s *TodoService) SendTodoMessage(item *models.TodoMailItem) error {

	var messageBody, subject string

	switch item.TodoEventType {
	case models.TodoEventTypeCreateTodo:
		messageBody = fmt.Sprintf(models.EmailBodyCreateTodo, item.Description, item.AssigneeName)
		subject = models.EmailSubjectCreateTodo

	case models.TodoEventTypeUpdateTodo:
		messageBody = fmt.Sprintf(models.EmailBodyUpdateTodo, item.Description, item.AssigneeName)
		subject = models.EmailSubjectUpdateTodo

	case models.TodoEventTypeDeleteTodo:
		messageBody = fmt.Sprintf(models.EmailBodyDeleteTodo, item.Description)
		subject = models.EmailSubjectDeleteTodo

	default:
		return app_errors.ErrIncorrectTodoEventType
	}

	s.logger.Info().Msgf("[SendUserMessage] sending %s message", subject)
	err := s.smtpClient.Send(item.Receivers, subject, messageBody)
	if err != nil {
		return fmt.Errorf("[SendTodoMessage]: send mssg: %w", err)
	}

	return nil
}
