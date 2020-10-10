package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/juanfer2/api-rest-go/controllers"
)

// "github.com/juanfer2/api-rest-go/controllers"
func helloWorld(c *fiber.Ctx) error {
	msg := fmt.Sprintf("Hello, %s 👋!", "World")
	return c.SendString(msg) // => Hello john 👋!
}

func SetupRoutes(app *fiber.App) {
	// Home
	app.Get("/", helloWorld)

	// Tasks
	app.Get("/tasks", controllers.GetTasks)
	app.Post("/tasks", controllers.CreateTasks)
	app.Get("/tasks/:id", controllers.GetTasksById)
	app.Put("/tasks/:id", controllers.UpdateTaskById)
	app.Delete("/tasks/:id", controllers.DeleteTaskById)

	// Users
	app.Get("/users", controllers.GetUsers)
	app.Post("/users", controllers.CreateUser)
}
