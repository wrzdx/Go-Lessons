package service

import (
	"context"
	"restapi/internal/core/domain"
	"restapi/internal/repository"
)

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *taskService {
	return &taskService{
		repo: repo,
	}
}

func (ts *taskService) Create(ctx context.Context, task TaskSnapshot) (TaskSnapshot, error) {
	newTask, err := domain.NewTask(task.GetTitle(), task.GetDescription())
	if err != nil {
		return nil, nil
	}
	dbModel := toRepoModel(newTask)

	if err := ts.repo.Create(ctx, dbModel); err != nil {
		return nil, err
	}

	return newTask, nil
}

func (ts *taskService) List(ctx context.Context) ([]TaskSnapshot, error) {
	return ts.List(ctx)
}
func (ts *taskService) ListUncompleted(ctx context.Context) ([]TaskSnapshot, error) {
	return ts.ListUncompleted(ctx)
}
func (ts *taskService) Update(ctx context.Context, title string, task repository.TaskPatch) (TaskSnapshot, error) {
	return ts.Update(ctx, title, task)
}
func (ts *taskService) Delete(ctx context.Context, title string) error {
	return ts.Delete(ctx, title)
}

func (ts *taskService) Get(ctx context.Context, title string) (TaskSnapshot, error) {
	return ts.repo.Get(ctx, title)
}
