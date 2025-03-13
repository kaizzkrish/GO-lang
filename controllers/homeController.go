package controllers

import (
	"fmt"
	"go-tutorial/database"
	"go-tutorial/models"

	"github.com/gofiber/fiber/v2"
)

func GetHome(c *fiber.Ctx) error {
	fmt.Println("Home Controller")
	db := database.DB
	var users []models.User
	db.Find(&users)
	return c.Render("index", fiber.Map{
		"Title": "Home",
		"Users": users,
	})
}
