package main

import (
	"fmt"
	"go-tutorial/database"
	"go-tutorial/models"
	"go-tutorial/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

func main() {
	fmt.Println("Program Started")
	// Initialize the HTML template engine
	engine := html.New("./views", ".html")

	// Create a new Fiber app with the HTML template engine
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layout",
	})

	// Connect to the database
	database.Connect()

	// Migrate the User model
	models.Migrate(database.DB)

	// Routes for rendering views
	routes.SetupRoutes(app)

	app.Use(logger.New())  // Logs each request
	app.Use(recover.New()) // Recovers from panics

	// Start the server
	app.Listen(":8080")
	fmt.Println("Server Started")
}
