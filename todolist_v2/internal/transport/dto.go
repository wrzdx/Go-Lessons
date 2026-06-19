package transport

import (
	"encoding/json"
	"restapi/internal/core"
	"time"
)

type CompleteTaskDTO struct {
	Complete bool
}

type TaskDTO struct {
	Title       string
	Description string
}

type PatchDRO struct {
	
}

func (t TaskDTO) ValidateForCreate() error {
	if t.Title == "" {
		return core.ErrEmptyTitle
	}

	return nil
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}

	return string(b)
}
