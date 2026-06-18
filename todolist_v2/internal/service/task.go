package service

import (
	"context"
	"maps"
	"restapi/internal/domain"
	"sync"
	"time"
)

type UpdateTask struct {
	Title       *string
	Description *string
	Completed   *bool

	CreatedAt   *time.Time
	CompletedAt *time.Time
}

type TaskRepository interface {
	Create(ctx context.Context, task domain.Task) (domain.Task, error)
	List(ctx context.Context) ([]domain.Task, error)
	ListUncompleted(ctx context.Context) ([]domain.Task, error)
	Get(ctx context.Context, title string) (domain.Task, error)
	Update(ctx context.Context, title string, task UpdateTask) (domain.Task, error)
	Delete(ctx context.Context, title string) error
}

type TaskService struct {
	tasks map[string]domain.Task
	mtx   sync.RWMutex
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks: make(map[string]domain.Task),
	}
}

func (l *TaskService) AddTask(task domain.Task) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if _, ok := l.tasks[task.Title]; ok {
		return domain.ErrTaskAlreadyExists
	}

	l.tasks[task.Title] = task

	return nil
}

func (l *TaskService) GetTask(title string) (domain.Task, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	task, ok := l.tasks[title]
	if !ok {
		return domain.Task{}, domain.ErrTaskNotFound
	}

	return task, nil
}

func (l *TaskService) ListTasks() map[string]domain.Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	tmp := make(map[string]domain.Task, len(l.tasks))

	maps.Copy(tmp, l.tasks)

	return tmp
}

func (l *TaskService) ListUncompletedTasks() map[string]domain.Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	uncompletedTasks := make(map[string]domain.Task)

	for title, task := range l.tasks {
		if !task.Completed {
			uncompletedTasks[title] = task
		}
	}

	return uncompletedTasks
}

func (l *TaskService) CompleteTask(title string) (domain.Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	task, ok := l.tasks[title]
	if !ok {
		return domain.Task{}, domain.ErrTaskNotFound
	}

	task.Complete()

	l.tasks[title] = task

	return task, nil
}

func (l *TaskService) UncompleteTask(title string) (domain.Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	task, ok := l.tasks[title]
	if !ok {
		return domain.Task{}, domain.ErrTaskNotFound
	}

	task.Uncomplete()

	l.tasks[title] = task

	return task, nil
}

func (l *TaskService) DeleteTask(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	_, ok := l.tasks[title]
	if !ok {
		return domain.ErrTaskNotFound
	}

	delete(l.tasks, title)

	return nil
}
