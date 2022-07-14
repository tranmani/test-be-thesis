package main

import (
	"testgo/test"
	"github.com/gofiber/fiber/v2"
)

func status(c *fiber.Ctx) error {
	return c.SendString("Server is running! Send your request")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", test.GetAll)
}

func main() {
	app := fiber.New()
	
	setupRoutes(app)
	app.Listen(":7000")
}