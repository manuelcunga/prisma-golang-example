package controllers

import (
	"net/http"
	"prisma-golang/dtos"
	services "prisma-golang/services/todo"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todoService services.TodoService
}

func NewTodoController(service services.TodoService) *TodoController {
	return &TodoController{todoService: service}
}

func (ctrl *TodoController) Create(ctx *gin.Context) error {
	var todoInput dtos.CreateTodoInput

	if err := ctx.ShouldBindJSON(&todoInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	todoOutput, err := ctrl.todoService.Create(&todoInput)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil
	}

	ctx.JSON(http.StatusCreated, gin.H{"todo": todoOutput})

	return nil

}

func (ctrl *TodoController) FindOne(ctx *gin.Context) error {
	todoID := ctx.Param("todo_id")

	todoOutput, err := ctrl.todoService.FindOne(todoID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"todo": "todo not found"})
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{"response": todoOutput})

	return nil
}

func (t *TodoController) ListAll(ctx *gin.Context) error {
	todos, err := t.todoService.ListAll()

	if err != nil {
		return err
	}

	var todosOutput []dtos.TodoOutput

	for _, todo := range *todos {
		todosOutput = append(todosOutput, dtos.TodoOutput{
			ID:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			CreatedAt:   todo.CreatedAt,
			UpdatedAt:   todo.UpdatedAt,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"todos": todosOutput})

	return nil
}
