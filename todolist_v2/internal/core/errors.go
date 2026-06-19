package core

import "errors"

var ErrTaskNotFound = errors.New("task not found")
var ErrTaskAlreadyExists = errors.New("task already exists")
var ErrEmptyTitle = errors.New("task title cannot be empty")
var ErrEmptyBody = errors.New("failed to read body")
var ErrInvalidCompleted = errors.New("invalid completed value, must be true or false")