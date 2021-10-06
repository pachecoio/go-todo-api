package todo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

func Register(router fiber.Router, database *gorm.DB) {
	database.AutoMigrate(&Todo{})
	todoRepository := NewTodoRepository(database)
	todoHandler := NewTodoHandler(todoRepository)

	movieRouter := router.Group("/todo")
	movieRouter.Get("/", todoHandler.GetAll)
	movieRouter.Get("/:id", todoHandler.Get)
	movieRouter.Put("/:id", todoHandler.Update)
	movieRouter.Post("/", todoHandler.Create)
	movieRouter.Delete("/:id", todoHandler.Delete)
}
