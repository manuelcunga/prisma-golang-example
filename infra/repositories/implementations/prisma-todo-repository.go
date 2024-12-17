package repository

import (
	"prisma-golang/dtos"
	"prisma-golang/infra/database"
	repository "prisma-golang/infra/repositories/interfaces"
	"prisma-golang/prisma/db"
)

type TodoImplementRepository struct {
	connection database.PrismaDB
}

func NewTodoRepository(conn database.PrismaDB) repository.ITodoRepository {
	return &TodoImplementRepository{connection: conn}
}

func (t *TodoImplementRepository) Create(data dtos.CreateTodoInput) (*dtos.TodoOutput, error) {

	todo, err := t.connection.Client.Todo.CreateOne(
		db.Todo.Title.Set(data.Title),
		db.Todo.Description.Set(data.Description),
	).Exec(t.connection.Context)

	if err != nil {
		return nil, err
	}

	return &dtos.TodoOutput{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}, nil
}

func (t *TodoImplementRepository) FindOne(todoID string) (*dtos.TodoOutput, error) {
	todo, err := t.connection.Client.Todo.FindUnique(
		db.Todo.ID.Equals(todoID),
	).Exec(t.connection.Context)

	if err != nil {
		return nil, err
	}

	todoOutput := &dtos.TodoOutput{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}

	return todoOutput, nil
}

func (t *TodoImplementRepository) ListAll() ([]dtos.TodoOutput, error) {
	todos, err := t.connection.Client.Todo.FindMany().Exec(t.connection.Context)

	if err != nil {
		return nil, err
	}

	var todosOutput []dtos.TodoOutput
	for _, todo := range todos {
		todosOutput = append(todosOutput, dtos.TodoOutput{
			ID:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			CreatedAt:   todo.CreatedAt,
			UpdatedAt:   todo.UpdatedAt,
		})
	}

	return todosOutput, nil
}
