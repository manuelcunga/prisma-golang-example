package repository

import "prisma-golang/dtos"

type ITodoRepository interface {
	Create(data dtos.CreateTodoInput) (*dtos.TodoOutput, error)
	ListAll() ([]dtos.TodoOutput, error)
	FindOne(todoID string) (*dtos.TodoOutput, error)
}
