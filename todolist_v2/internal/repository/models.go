package repository

import "time"

type TaskModel struct {
	Title       string
	Description string
	Completed   bool

	CreatedAt   time.Time
	CompletedAt *time.Time
}

func (m TaskModel) GetTitle() string           { return m.Title }
func (m TaskModel) GetDescription() string     { return m.Description }
func (m TaskModel) GetCompleted() bool         { return m.Completed }
func (m TaskModel) GetCreatedAt() time.Time    { return m.CreatedAt }
func (m TaskModel) GetCompletedAt() *time.Time { return m.CompletedAt }
