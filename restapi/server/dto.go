package server

import (
	"encoding/json"
	"errors"
	"time"
)

type BookDTO struct {
	Title  string
	Author string
	Size   int
}

type ReadBookDTO struct {
	Read *bool
}

func (b BookDTO) ValidateForCreate() error {
	if b.Title == "" {
		return errors.New("title is empty")
	}

	if b.Author == "" {
		return errors.New("author is empty")
	}

	if b.Size <= 0 {
		return errors.New("size is not valid")
	}

	return nil
}

func (r ReadBookDTO) ValidateForRead() error {
	if r.Read == nil {
		return errors.New("read field is required, bool type")
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
