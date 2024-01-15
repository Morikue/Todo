package models

const (
	TodoEventTypeCreateTodo = "create_todo"
	TodoEventTypeUpdateTodo = "update_todo"
	TodoEventTypeDeleteTodo = "delete_todo"
)

type TodoMailItem struct {
	TodoEventType string   `json:"todo_event_type"`
	Receivers     []string `json:"receivers"`
	AssigneeName  string   `json:"assignee_name"`
	Description   string   `json:"description"`
}
