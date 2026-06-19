package service

import (
	"restapi/internal/core"
	"strings"
	"time"
)

type Task struct {
	title       string
	description string
	completed   bool

	createdAt   time.Time
	completedAt *time.Time
}

func fromSnapshot(snapshot TaskSnapshot) Task {
	return Task{
		title:       snapshot.GetTitle(),
		description: snapshot.GetDescription(),
		completed:   snapshot.GetCompleted(),
		createdAt:   snapshot.GetCreatedAt(),
		completedAt: snapshot.GetCompletedAt(),
	}
}

func NewTask(title string, description string) (Task, error) {
	title = strings.TrimSpace(title)
	description = strings.TrimSpace(description)
	if len(title) == 0 {
		return Task{}, core.ErrEmptyTitle
	}
	return Task{
		title:       title,
		description: description,
		completed:   false,

		createdAt:   time.Now(),
		completedAt: nil,
	}, nil
}

func (t Task) GetTitle() string           { return t.title }
func (t Task) GetDescription() string     { return t.description }
func (t Task) GetCompleted() bool         { return t.completed }
func (t Task) GetCompletedAt() *time.Time { return t.completedAt }
func (t Task) GetCreatedAt() time.Time    { return t.createdAt }

func (t *Task) SetTitle(title string) error {
	title = strings.TrimSpace(title)
	if len(title) == 0 {
		return core.ErrEmptyTitle
	}
	t.title = title
	return nil
}

func (t *Task) SetDescription(description string) {
	t.description = description
}

func (t *Task) SetCompleted(completed bool) {
    if t.completed == completed {
        return
    }

    if completed {
        completeTime := time.Now()
        t.completed = true
        t.completedAt = &completeTime
    } else {
        t.completed = false
        t.completedAt = nil
    }
}
