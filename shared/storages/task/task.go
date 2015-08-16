package task

import (
	"time"
	"errors"
)

type Task struct {
	Id          int64 `json:"id"`
	Title       string `sql:"default: not null", json:"title"`
	Description *string `sql:"default: null", json:"description"`
	Priority    *int64 `sql:"default: null", json:"priority"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	CompletedAt *time.Time `json:"completedAt"`
	IsDeleted   bool `json:"isDeleted"`
	IsCompleted bool `json:"isCompleted"`
}

func(t *Task) AfterFind() error {
	return nil
}

var (
	ErrTaskNotFound = errors.New("Task not found")
)