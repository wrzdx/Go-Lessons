package repository

import (
	"context"
	"database/sql"
	"errors"
	"restapi/internal/core"
	"restapi/internal/service"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type postgresRepository struct {
	conn *pgx.Conn
}

func NewPostgres(conn *pgx.Conn) *postgresRepository {
	return &postgresRepository{
		conn: conn,
	}
}

func (db *postgresRepository) Create(ctx context.Context, t service.TaskSnapshot) error {
	query := `
		INSERT INTO tasks 
		(title, description, completed, created_at, completed_at) 
		VALUES 
		($1, $2, $3, $4, $5);
	`
	_, err := db.conn.Exec(
		ctx,
		query,
		t.GetTitle(),
		t.GetDescription(),
		t.GetCompleted(),
		t.GetCreatedAt(),
		t.GetCompletedAt(),
	)

	if err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
			if pgErr.Code == "23505" { // unique_violation
				return core.ErrTaskAlreadyExists
			}
		}
		return err
	}
	return nil
}
func (db *postgresRepository) List(ctx context.Context, completedFilter *bool) ([]service.TaskSnapshot, error) {
	var query string
	var args []any

	if completedFilter == nil {
		query = "SELECT * FROM tasks;"
	} else {
		query ="SELECT * FROM tasks WHERE completed = $1;"
		args = append(args, *completedFilter)
	}
	rows, err := db.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := []service.TaskSnapshot{}
	for rows.Next() {
		var t TaskModel
		if err := rows.Scan(&t.Title,
			&t.Description,
			&t.Completed,
			&t.CreatedAt,
			&t.CompletedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}


func (db *postgresRepository) Get(ctx context.Context, title string) (service.TaskSnapshot, error) {
	var t TaskModel
	query := `SELECT * FROM tasks WHERE title=$1;`
	err := db.conn.QueryRow(
		ctx,
		query,
		title,
	).Scan(
		&t.Title,
		&t.Description,
		&t.Completed,
		&t.CreatedAt,
		&t.CompletedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return TaskModel{}, core.ErrTaskNotFound
		}
		return TaskModel{}, err
	}
	return t, nil
}

func (db *postgresRepository) Update(ctx context.Context, title string, t service.TaskSnapshot) (service.TaskSnapshot, error) {
	var updated TaskModel
	query := `
		UPDATE tasks 
		SET title=COALESCE($1, title), 
			description=COALESCE($2, description), 
			completed=$3,  
			completed_at=$4
		WHERE title=$5
		RETURNING *;
	`
	err := db.conn.QueryRow(
		ctx,
		query,
		t.GetTitle(),
		t.GetDescription(),
		t.GetCompleted(),
		t.GetCompletedAt(),
		title,
	).Scan(
		&updated.Title,
		&updated.Description,
		&updated.Completed,
		&updated.CreatedAt,
		&updated.CompletedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return TaskModel{}, core.ErrTaskNotFound
		}
		return TaskModel{}, err
	}
	return updated, nil
}

func (db *postgresRepository) Delete(ctx context.Context, title string) error {
	query := `DELETE FROM tasks WHERE title = $1;`

	result, err := db.conn.Exec(ctx, query, title)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return core.ErrTaskNotFound
	}

	return nil
}
