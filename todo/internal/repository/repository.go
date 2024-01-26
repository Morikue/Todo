package repository

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/opentracing/opentracing-go"
	"todo/internal/models"
	"todo/pkg/ctxutil"
)

type TodoRepository struct {
	conn *pgxpool.Pool
}

func NewTodoRepository(conn *pgxpool.Pool) *TodoRepository {
	return &TodoRepository{conn}
}

func (r *TodoRepository) CreateToDo(ctx context.Context, newTodo *models.TodoDAO) (*models.TodoDAO, error) {
	ctx = ctxutil.SetRequestIdFromGrpcToContext(ctx)

	span, ctx := opentracing.StartSpanFromContext(ctx, "repository.CreateToDo")
	defer span.Finish()

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
	ctx = ctxutil.SetRequestIdFromGrpcToContext(ctx)

	span, ctx := opentracing.StartSpanFromContext(ctx, "repository.UpdateToDo")
	defer span.Finish()

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
	ctx = ctxutil.SetRequestIdFromGrpcToContext(ctx)

	span, ctx := opentracing.StartSpanFromContext(ctx, "repository.GetToDos")
	defer span.Finish()

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
	ctx = ctxutil.SetRequestIdFromGrpcToContext(ctx)

	span, ctx := opentracing.StartSpanFromContext(ctx, "repository.GetToDo")
	defer span.Finish()

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
	ctx = ctxutil.SetRequestIdFromGrpcToContext(ctx)

	span, ctx := opentracing.StartSpanFromContext(ctx, "repository.DeleteToDo")
	defer span.Finish()

	sql := `
        DELETE FROM 
		    todos
        WHERE 
            id = $1
    `
	_, err := r.conn.Exec(ctx, sql, todoID)
	return err
}
