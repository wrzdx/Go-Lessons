package repository

import (
	"context"
)

type TaskPatch interface {
	GetTitle() *string
	GetDescription() *string
	GetCompleted() *bool
}

type TaskRepository interface {
	Create(ctx context.Context, task TaskModel) error
	List(ctx context.Context) ([]TaskModel, error)
	ListUncompleted(ctx context.Context) ([]TaskModel, error)
	Get(ctx context.Context, title string) (TaskModel, error)
	Update(ctx context.Context, title string, task TaskPatch) (TaskModel, error)
	Delete(ctx context.Context, title string) error
}
