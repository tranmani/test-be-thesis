package test

import (
	"testgo/database"
	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {
	result, err := database.GetAll()
	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.Status(200).JSON(result)
}