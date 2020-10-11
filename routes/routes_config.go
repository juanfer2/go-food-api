package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/juanfer2/api-rest-go/constants"
	"github.com/juanfer2/api-rest-go/controllers"
)

// "github.com/juanfer2/api-rest-go/controllers"
func helloWorld(c *fiber.Ctx) error {
	msg := fmt.Sprintf("Hello, %s 👋!", "World")
	return c.SendString(msg) // => Hello john 👋!
}

func skipRoutes(app *fiber.App) {
	app.Post("/users", controllers.CreateUser)

	// Food
	app.Get("/foods", controllers.GetFoods)
	app.Get("/foods/:id", controllers.GetFoodById)

	// TypeFood
	app.Get("/type_foods", controllers.GetTypeFoods)
	app.Get("/type_foods/:id", controllers.GetTypeFoodById)
}

func protectedRoutes(app *fiber.App) {
	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(constants.SecretKey),
	}))

	// Users
	app.Get("/users", controllers.GetUsers)

	// Order
	app.Get("/orders", controllers.GetOrders)
	app.Post("/orders", controllers.CreateOrder)
	app.Get("/orders/:id", controllers.GetOrderById)
	app.Put("/orders/:id", controllers.UpdateOrderById)
	app.Delete("/orders/:id", controllers.DeleteOrderById)

	// Food
	app.Post("/foods", controllers.CreateFood)
	app.Put("/foods/:id", controllers.UpdateFoodById)
	app.Delete("/foods/:id", controllers.DeleteFoodById)

	// TypeFood
	app.Post("/type_foods", controllers.CreateTypeFood)
	app.Put("/type_foods/:id", controllers.UpdateTypeFoodById)
	app.Delete("/type_foods/:id", controllers.DeleteTypeFoodById)
}

func SetupRoutes(app *fiber.App) {
	// Home
	app.Get("/", helloWorld)
	skipRoutes(app)
	protectedRoutes(app)
}
