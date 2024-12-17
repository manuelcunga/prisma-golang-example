package dependency

import (
	"prisma-golang/infra/database"
	"prisma-golang/infra/http/controllers"
	repository "prisma-golang/infra/repositories/implementations"
	services "prisma-golang/services/todo"
)

func TodoDependency(connection database.PrismaDB) *controllers.TodoController {
	todoRepo := repository.NewTodoRepository(connection)
	todoService := services.NewCreateTodoService(todoRepo)
	todoController := controllers.NewTodoController(*todoService)
	return todoController
}
