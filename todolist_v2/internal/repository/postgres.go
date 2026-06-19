package repository

import (
	"context"
	"database/sql"
	"errors"
	"restapi/internal/core"
	"restapi/internal/service"

	"github.com/jackc/pgx/v5"
)

type postgresRepository struct {
	conn *pgx.Conn
}

func NewPostgres(conn *pgx.Conn) *postgresRepository {
	return &postgresRepository{
		conn: conn,
	}
}

func (db *postgresRepository) Create(ctx context.Context, t TaskModel) (TaskModel, error) {
	query := `
		INSERT INTO tasks 
		(title, description, completed, created_at, completed_at) 
		VALUES 
		($1, $2, $3, $4, $5)
		RETURNING *;
	`
	err := db.conn.QueryRow(
		ctx,
		query,
		t.Title,
		t.Description,
		t.Completed,
		t.CreatedAt,
		t.CompletedAt,
	).Scan(
		&t.Title,
		&t.Description,
		&t.Completed,
		&t.CreatedAt,
		&t.CompletedAt,
	)

	if err != nil {
		return TaskModel{}, err
	}
	return t, nil
}
func (db *postgresRepository) List(ctx context.Context) ([]TaskModel, error) {
	query := `SELECT * FROM tasks;`
	rows, err := db.conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := []TaskModel{}
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

func (db *postgresRepository) ListUncompleted(ctx context.Context) ([]TaskModel, error) {
	query := `SELECT * FROM tasks WHERE completed=false;`
	rows, err := db.conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := []TaskModel{}
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

func (db *postgresRepository) Get(ctx context.Context, title string) (TaskModel, error) {
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

func (db *postgresRepository) Update(ctx context.Context, title string, t service.TaskPatch) (TaskModel, error) {
	var updated TaskModel
	query := `
		UPDATE tasks 
		SET title=COALESCE($1, title), 
			description=COALESCE($2, description), 
			completed=COALESCE($3, completed), 
			created_at=COALESCE($4, created_at), 
			completed_at=COALESCE($5, completed_at)
		WHERE title=$6
		RETURNING *;
	`
	err := db.conn.QueryRow(
		ctx,
		query,
		t.GetTitle(),
		t.GetDescription(),
		t.GetCompleted(),
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
