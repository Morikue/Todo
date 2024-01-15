package models

const (
	TodoEventTypeCreateTodo = "create_todo"
	TodoEventTypeUpdateTodo = "update_todo"
	TodoEventTypeDeleteTodo = "delete_todo"
)

const (
	EmailSubjectCreateTodo = "A new TODO has been created"
	EmailSubjectUpdateTodo = "Your TODO has been changed"
	EmailSubjectDeleteTodo = "Your TODO has been deleted"
)

const (
	EmailBodyCreateTodo = `
		<!DOCTYPE html>
		<html>
		<body>
		
			<h2>A new TODO has been created!</h2> 

			<div> 
			  <h3>Description:</h3> 
			  <p>%s</p> 
			</div> 
    
			<div> 
			  <h3>Assignee:</h3> 
			  <p>%s</p> 
			</div> 

		</body>
		</html>
	`
	EmailBodyUpdateTodo = `
		<!DOCTYPE html>
		<html>
		<body>
		
			<h2>Your TODO has been changed!</h2> 

			<div> 
			  <h3>Description:</h3> 
			  <p>%s</p> 
			</div> 
    
			<div> 
			  <h3>Assignee:</h3> 
			  <p>%s</p> 
			</div> 

		</body>
		</html>
	`

	EmailBodyDeleteTodo = `
		<!DOCTYPE html>
		<html>
		<body>
		
			<h2>Your TODO has been deleted!</h2> 

			<div> 
			  <h3>Description:</h3> 
			  <p>%s</p> 
			</div>

		</body>
		</html>
	`
)

type TodoMailItem struct {
	TodoEventType string   `json:"todo_event_type"`
	Receivers     []string `json:"receivers"`
	AssigneeName  string   `json:"assignee_name"`
	Description   string   `json:"description"`
}
