package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetHome(c *fiber.Ctx) error {
	fmt.Println("Home Controller")
	return c.Render("index", fiber.Map{
		"Title": "Home",
	})
}
