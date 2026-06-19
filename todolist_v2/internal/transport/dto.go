package transport

import (
	"encoding/json"
	"restapi/internal/core"
	"restapi/internal/service"
	"time"
)

type taskResponse struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

func snapshotToResponse(snapshot service.TaskSnapshot) taskResponse {
	return taskResponse{
		Title:       snapshot.GetTitle(),
		Description: snapshot.GetDescription(),
		Completed:   snapshot.GetCompleted(),
		CreatedAt:   snapshot.GetCreatedAt(),
		CompletedAt: snapshot.GetCompletedAt(),
	}
}

type taskInput struct {
	Title       string
	Description string
}

func (t taskInput) GetTitle() string       { return t.Title }
func (t taskInput) GetDescription() string { return t.Description }

func (t taskInput) ValidateForCreate() error {
	if t.Title == "" {
		return core.ErrEmptyTitle
	}

	return nil
}

type patchDTO struct {
	Title       *string
	Description *string
	Completed   *bool
}

func (p patchDTO) GetTitle() *string       { return p.Title }
func (p patchDTO) GetDescription() *string { return p.Description }
func (p patchDTO) GetCompleted() *bool     { return p.Completed }

type errorDTO struct {
	Message string
	Time    time.Time
}

func (e errorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}

	return string(b)
}
