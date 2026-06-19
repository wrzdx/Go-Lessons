package core

import "errors"

var ErrTaskNotFound = errors.New("task not found")
var ErrTaskAlreadyExists = errors.New("task already exists")
var ErrEmptyTitle = errors.New("task title cannot be empty")
