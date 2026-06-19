package service

import (
	"context"
	"restapi/internal/repository"
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

type TaskService interface {
	Create(ctx context.Context, task TaskInput) (TaskSnapshot, error)
	List(ctx context.Context) ([]TaskSnapshot, error)
	ListUncompleted(ctx context.Context) ([]TaskSnapshot, error)
	Get(ctx context.Context, title string) (TaskSnapshot, error)
	Update(ctx context.Context, title string, task repository.TaskPatch) (TaskSnapshot, error)
	Delete(ctx context.Context, title string) error
}

func toRepoModel(task TaskSnapshot) repository.TaskModel {
	return repository.TaskModel{
		Title:       task.GetTitle(),
		Description: task.GetDescription(),
		Completed:   task.GetCompleted(),
		CreatedAt:   task.GetCreatedAt(),
		CompletedAt: task.GetCompletedAt(),
	}
}
