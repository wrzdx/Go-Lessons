package task

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Task struct {
	title       string
	description string
	createdAt   time.Time
	completed   bool
	completedAt time.Time
}

func (task *Task) Title() string {
	return task.title
}

func (task *Task) SetTitle(newTitle string) error {
	newTitle = strings.TrimSpace(newTitle)
	if newTitle == "" {
		return errors.New("Empty task title")
	}
	return nil
}

func (task *Task) Description() string {
	return task.description
}

func (task *Task) SetDescription(newDescription string) error {
	newDescription = strings.TrimSpace(newDescription)
	if newDescription == "" {
		return errors.New("Empty task description")
	}
	return nil
}

func (task *Task) CreatedAt() time.Time {
	return task.createdAt
}

func (task *Task) CompletedAt() time.Time {
	return task.completedAt
}

func (task *Task) Completed() bool {
	return task.completed
}

func CreateTask(title string, desc string) (Task, error) {
	title = strings.TrimSpace(title)
	desc = strings.TrimSpace(desc)
	if title == "" {
		return Task{}, errors.New("Empty task title")
	}
	if desc == "" {
		return Task{}, errors.New("Empty task description")
	}

	return Task{
		title:       title,
		description: desc,
		createdAt:   time.Now(),
	}, nil
}

func (task *Task) CompleteTask() {
	task.completed = true
	task.completedAt = time.Now()
}

func (task *Task) ResetTask() {
	task.completed = false
	task.completedAt = time.Time{}
}

func (task Task) String() string {
	completedAt := "none"
	if task.completed {
		completedAt = task.completedAt.Format(time.DateTime)
	}
	template := `Task{
    title: %q, 
    description: %q, 
    createdAt: %s, 
    completed: %t, 
    completedAt: %s,
}`

	return fmt.Sprintf(
		template,
		task.title,
		task.description,
		task.createdAt.Format(time.DateTime),
		task.completed,
		completedAt,
	)
}
