package controllers

import (
	"go-tutorial/database"
	"go-tutorial/models"

	"github.com/gofiber/fiber/v2"
)

func GetUsersJSON(c *fiber.Ctx) error {
	db := database.DB
	var users []models.User
	db.Find(&users)
	return c.JSON(users)
}

func GetUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []models.User
	db.Find(&users)
	return c.Render("view_user", fiber.Map{
		"Title": "User List", // Add this line
		"Users": users,
	})
}

func AddUser(c *fiber.Ctx) error {
	db := database.DB
	user := new(models.User)

	// Handle form and JSON submissions
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Save the user to the database
	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not save user"})
	}

	// Redirect after adding a user (if using a form)
	return c.Redirect("/users")
}
