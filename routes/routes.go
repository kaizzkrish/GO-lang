package routes

import (
	"fmt"
	"go-tutorial/controllers"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up all the routes for the application
func SetupRoutes(app *fiber.App) {

	fmt.Println("Routes")
	app.Get("/", controllers.GetHome)
	app.Get("/users", controllers.GetUsers)
	app.Get("/api/users", controllers.GetUsersJSON)

	// ✅ Allow POST requests to add a user
	app.Post("/users", controllers.AddUser)
}
