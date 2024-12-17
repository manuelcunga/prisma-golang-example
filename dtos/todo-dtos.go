package dtos

import (
	"time"
)

type CreateTodoInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoOutput struct {
	ID          string
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
