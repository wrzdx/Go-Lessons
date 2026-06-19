package domain

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

func (t *Task) SetDescription(title string) {
	t.title = title
}

func (t *Task) SetCompleted(completed bool) {
	if completed {
		completeTime := time.Now()
		t.completed = true
		t.completedAt = &completeTime
	} else {
		t.completed = false
		t.completedAt = nil
	}
}
