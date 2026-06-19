package service

import (
	"context"
)

type taskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *taskService {
	return &taskService{
		repo: repo,
	}
}

func (ts *taskService) Create(ctx context.Context, task TaskInput) (TaskSnapshot, error) {
	newTask, err := NewTask(task.GetTitle(), task.GetDescription())
	if err != nil {
		return nil, nil
	}

	if err := ts.repo.Create(ctx, newTask); err != nil {
		return nil, err
	}

	return newTask, nil
}

func (ts *taskService) List(ctx context.Context, completedFilter *bool) ([]TaskSnapshot, error) {
	return ts.repo.List(ctx, completedFilter)
}
func (ts *taskService) Update(ctx context.Context, title string, patch TaskPatch) (TaskSnapshot, error) {
	snapshot, err := ts.repo.Get(ctx, title)
	if err != nil {
		return nil, err
	}
	currentTask := fromSnapshot(snapshot)
	if patch.GetTitle() != nil {
		if err := currentTask.SetTitle(*patch.GetTitle()); err != nil {
			return Task{}, err
		}
	}
	if patch.GetDescription() != nil {
		currentTask.SetDescription(*patch.GetDescription())
	}
	if patch.GetCompleted() != nil {
		currentTask.SetCompleted(*patch.GetCompleted())
	}

	return ts.repo.Update(ctx, title, currentTask)
}
func (ts *taskService) Delete(ctx context.Context, title string) error {
	return ts.repo.Delete(ctx, title)
}

func (ts *taskService) Get(ctx context.Context, title string) (TaskSnapshot, error) {
	return ts.repo.Get(ctx, title)
}
