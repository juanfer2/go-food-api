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

	// Order
	app.Get("/orders", controllers.GetOrders)
	app.Post("/orders", controllers.CreateOrder)
	app.Get("/orders/:id", controllers.GetOrderById)
	app.Put("/orders/:id", controllers.UpdateOrderById)
	app.Delete("/orders/:id", controllers.DeleteOrderById)

	// Food
	app.Get("/foods", controllers.GetFoods)
	app.Post("/foods", controllers.CreateFood)
	app.Get("/foods/:id", controllers.GetFoodById)
	app.Put("/foods/:id", controllers.UpdateFoodById)
	app.Delete("/foods/:id", controllers.DeleteFoodById)

	// TypeFood
	app.Get("/type_foods", controllers.GetTypeFoods)
	app.Post("/type_foods", controllers.CreateTypeFood)
	app.Get("/type_foods/:id", controllers.GetTypeFoodById)
	app.Put("/type_foods/:id", controllers.UpdateTypeFoodById)
	app.Delete("/type_foods/:id", controllers.DeleteTypeFoodById)
}
