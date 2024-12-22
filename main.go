package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shimodii/ecommerce-backend/repository"
	"github.com/shimodii/ecommerce-backend/route"
)

func main() {
	app := fiber.New()

	//calling repository InitDatabase to initialize database connection
	repository.OpenDatabase()

	route.SetupRoutes(app)

	app.Listen(":3344")
}
