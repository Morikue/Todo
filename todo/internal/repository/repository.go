package repository

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"todo/internal/models"
)

type TodoRepository struct {
	conn *pgxpool.Pool
}

func NewTodoRepository(conn *pgxpool.Pool) *TodoRepository {
	return &TodoRepository{conn}
}

func (r *TodoRepository) CreateToDo(ctx context.Context, newTodo *models.TodoDAO) (*models.TodoDAO, error) {
	var todoId uuid.UUID

	sql := `
        INSERT INTO 
			todos (
			    id,
				created_by, 
				assignee, 
				description,
				created_at,
				updated_at
			)
        VALUES 
			($1, $2, $3, $4, now(), now())
        RETURNING id
    `
	err := r.conn.QueryRow(ctx, sql, newTodo.ID, newTodo.CreatedBy, newTodo.Assignee, newTodo.Description).
		Scan(&todoId)
	if err != nil {
		return nil, err
	}

	newTodo.ID = todoId

	return newTodo, nil
}

func (r *TodoRepository) UpdateToDo(ctx context.Context, newTodo *models.TodoDAO) (*models.TodoDAO, error) {
	sql := `
	UPDATE
		todos
	SET
	    assignee = $2,
	    description = $3,
	    updated_at = now()
	WHERE 
	    id = $1
	`

	_, err := r.conn.Exec(ctx, sql, newTodo.ID, newTodo.Assignee, newTodo.Description)
	if err != nil {
		return nil, err
	}

	return newTodo, nil
}

func (r *TodoRepository) GetToDos(ctx context.Context, todos *models.GetTodosDTO) ([]models.TodoDAO, error) {
	var todosFromDb = make([]models.TodoDAO, 0)

	builder := squirrel.Select(
		"id",
		"created_by",
		"assignee",
		"description",
		"created_at",
		"updated_at",
	).
		From("todos").
		PlaceholderFormat(squirrel.Dollar)

	if todos.CreatedBy != 0 {
		builder = builder.Where(squirrel.Eq{"created_by": todos.CreatedBy})
	}

	if todos.Assignee != 0 {
		builder = builder.Where(squirrel.Eq{"assignee": todos.Assignee})
	}

	if !todos.DateFrom.IsZero() && !todos.DateTo.IsZero() {
		builder = builder.Where("created_at BETWEEN ? AND ?", todos.DateFrom, todos.DateTo)
	} else if !todos.DateFrom.IsZero() {
		builder = builder.Where("created_at >= ?", todos.DateFrom)
	} else if !todos.DateTo.IsZero() {
		builder = builder.Where("created_at <= ?", todos.DateTo)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("[GetToDos] build query: %w", err)
	}

	rows, err := r.conn.Query(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return nil, fmt.Errorf("[GetToDos] query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.TodoDAO
		err := rows.Scan(
			&todo.ID,
			&todo.CreatedBy,
			&todo.Assignee,
			&todo.Description,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		todosFromDb = append(todosFromDb, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("[GetToDos] scan: %w", err)
	}

	return todosFromDb, nil
}

func (r *TodoRepository) GetToDo(ctx context.Context, todoID uuid.UUID) (*models.TodoDAO, error) {
	var todo models.TodoDAO
	sql := `
        SELECT 
            id, 
			created_by, 
			assignee, 
			description,
			created_at,
			updated_at
        FROM 
            todos
        WHERE 
            id = $1
    `
	err := r.conn.QueryRow(ctx, sql, todoID).
		Scan(&todo.ID, &todo.CreatedBy, &todo.Assignee, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepository) DeleteToDo(ctx context.Context, todoID uuid.UUID) error {
	sql := `
        DELETE FROM 
		    todos
        WHERE 
            id = $1
    `
	_, err := r.conn.Exec(ctx, sql, todoID)
	return err
}
