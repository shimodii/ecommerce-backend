package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shimodii/ecommerce-backend/repository"
)

func main() {
	app := fiber.New()

	//calling repository InitDatabase to initialize database connection
	repository.InitDatabase()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("everything is up")
	})

	app.Listen(":3344")
}
