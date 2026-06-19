package service

import (
	"context"
	"time"
)

type TaskInput interface {
	GetTitle() string
	GetDescription() string
}

type TaskSnapshot interface {
	GetTitle() string
	GetDescription() string
	GetCompleted() bool
	GetCreatedAt() time.Time
	GetCompletedAt() *time.Time
}

type TaskPatch interface {
	GetTitle() *string
	GetDescription() *string
	GetCompleted() *bool
}
type TaskService interface {
	Create(ctx context.Context, task TaskInput) (TaskSnapshot, error)
	List(ctx context.Context) ([]TaskSnapshot, error)
	ListUncompleted(ctx context.Context) ([]TaskSnapshot, error)
	Get(ctx context.Context, title string) (TaskSnapshot, error)
	Update(ctx context.Context, title string, task TaskPatch) (TaskSnapshot, error)
	Delete(ctx context.Context, title string) error
}

type TaskRepository interface {
	Create(ctx context.Context, task TaskSnapshot) error
	List(ctx context.Context) ([]TaskSnapshot, error)
	ListUncompleted(ctx context.Context) ([]TaskSnapshot, error)
	Get(ctx context.Context, title string) (TaskSnapshot, error)
	Update(ctx context.Context, title string, task TaskPatch) (TaskSnapshot, error)
	Delete(ctx context.Context, title string) error
}
