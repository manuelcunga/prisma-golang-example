package router

import (
	"prisma-golang/infra/database"
	"prisma-golang/infra/http/dependency"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	Router *gin.Engine
}

func NewRouter(r *gin.Engine) *RouterConfig {
	return &RouterConfig{Router: r}
}

func (r *RouterConfig) RouterSetup(db database.PrismaDB) {
	todoController := dependency.TodoDependency(db)

	api := "/api/v1"

	r.Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	v1 := r.Router.Group(api)

	v1.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome todo API",
		})
	})

	v1.POST("/todos", func(ctx *gin.Context) {
		todoController.Create(ctx)
	})

	v1.GET("/todos", func(ctx *gin.Context) {
		todoController.ListAll(ctx)
	})

	v1.GET("/todos/:todo_id", func(ctx *gin.Context) {
		todoController.FindOne(ctx)
	})
}
