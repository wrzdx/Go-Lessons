package repository

import (
	"context"
	"restapi/internal/domain"
	"restapi/internal/service"

	"github.com/jackc/pgx/v5"
)

type PostgresRepository struct {
	conn *pgx.Conn
}

func NewPostgres(conn *pgx.Conn) *PostgresRepository {
	return &PostgresRepository{
		conn: conn,
	}
}

func (db *PostgresRepository) Create(ctx context.Context, t domain.Task) (domain.Task, error) {
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
		return domain.Task{}, err
	}
	return t, nil
}
func (db *PostgresRepository) List(ctx context.Context) ([]domain.Task, error) {
	query := `SELECT * FROM tasks;`
	rows, err := db.conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := []domain.Task{}
	for rows.Next() {
		var t domain.Task
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

func (db *PostgresRepository) ListUncompleted(ctx context.Context) ([]domain.Task, error) {
	query := `SELECT * FROM tasks WHERE completed=false;`
	rows, err := db.conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := []domain.Task{}
	for rows.Next() {
		var t domain.Task
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

func (db *PostgresRepository) Get(ctx context.Context, title string) (domain.Task, error) {
	var t domain.Task
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
		return domain.Task{}, err
	}
	return t, nil
}

func (db *PostgresRepository) Update(ctx context.Context, title string, t service.UpdateTask) (domain.Task, error) {
	var updated domain.Task
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
		t.Title,
		t.Description,
		t.Completed,
		t.CreatedAt,
		t.CompletedAt,
		title,
	).Scan(
		&updated.Title,
		&updated.Description,
		&updated.Completed,
		&updated.CreatedAt,
		&updated.CompletedAt,
	)

	if err != nil {
		return domain.Task{}, err
	}
	return updated, nil
}

func (db *PostgresRepository) Delete(ctx context.Context, title string) error {
    query := `DELETE FROM tasks WHERE title = $1;`
    
    result, err := db.conn.Exec(ctx, query, title)
    if err != nil {
        return err
    }
    
    if result.RowsAffected() == 0 {
        return domain.ErrTaskNotFound 
    }
    
    return nil
}
