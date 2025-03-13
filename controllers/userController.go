package controllers

import (
	"fmt"
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

	// Debugging: Print users in console
	fmt.Println("Users retrieved from database:", users)

	return c.Render("view_user", fiber.Map{
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

func EditUserPage(c *fiber.Ctx) error {
	id := c.Params("id") // Get ID from URL
	db := database.DB
	var user models.User

	// Find the user by ID
	if err := db.First(&user, id).Error; err != nil {
		fmt.Println("User not found for ID:", id)
		return c.Status(404).SendString("User not found")
	}

	// Debugging: Print retrieved user
	fmt.Println("Editing User:", user)

	// Pass user data to template
	return c.Render("edit_user", fiber.Map{
		"User": user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	db := database.DB
	id := c.Params("id") // Get user ID from URL

	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	// Parse form data
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Save updated user
	if err := db.Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not update user"})
	}

	return c.Redirect("/users") // Redirect back to user list
}
