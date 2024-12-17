package services

import (
	"fmt"
	"prisma-golang/dtos"
	repository "prisma-golang/infra/repositories/interfaces"
)

type TodoService struct {
	todoRepository repository.ITodoRepository
}

func NewCreateTodoService(todoRepo repository.ITodoRepository) *TodoService {
	return &TodoService{todoRepository: todoRepo}
}

func (t *TodoService) Create(input *dtos.CreateTodoInput) (*dtos.TodoOutput, error) {
	todo, err := t.todoRepository.Create(*input)
	if err != nil {
		return nil, err
	}

	fmt.Println("cadastrado", todo)

	return &dtos.TodoOutput{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}, nil
}

func (t *TodoService) FindOne(todoID string) (*dtos.TodoOutput, error) {
	todo, err := t.todoRepository.FindOne(todoID)
	if err != nil {
		return nil, err
	}

	todoOutput := &dtos.TodoOutput{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.CreatedAt,
	}

	return todoOutput, nil
}

func (t *TodoService) ListAll() (*[]dtos.TodoOutput, error) {
	todos, err := t.todoRepository.ListAll()

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

	return &todosOutput, nil

}
