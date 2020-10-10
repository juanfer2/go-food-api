package servers

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/arsmn/fiber-swagger/v2/example/docs"
	"github.com/gofiber/fiber/v2"

	"github.com/juanfer2/api-rest-go/databases"
	"github.com/juanfer2/api-rest-go/models"
	"github.com/juanfer2/api-rest-go/routes"
)

func StartServerApp() {
	db := databases.Conn()

	// Migrate the schema
	db.AutoMigrate(&models.Task{}, &models.User{})

	// logger, _ := zap.NewProduction()
	// defer logger.Sync()
	// logger.Warn(
	// 	"title",
	// Structured context as strongly typed Field values.
	//	zap.String("url", "url"),
	//	zap.Int("attempt", 3),
	// zap.Duration("backoff", time.Second),
	//)

	// db := databases.Conn()
	// log.Println(db.AutoMigrate(&models.Task{}).Error)
	// Viper will use this and not check any of the config paths.
	// .env - It will search for the .env file in the current directory

	app := fiber.New()
	app.Use("/api-doc", swagger.Handler) // default

	app.Use("/api-doc", swagger.New(swagger.Config{ // custom
		URL:         "./docs/doc.json",
		DeepLinking: true,
	}))
	routes.SetupRoutes(app)
	app.Listen(":4000")
}
