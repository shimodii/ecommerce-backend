package router

import "github.com/gofiber/fiber/v2"

func SetupRouter(app *fiber.App){
    // api for public apis
    api := app.Use("api")

    // admin for admin apis
    // admin := app.Use("admin")

    //healthcheck
    app.Get("/status", func(c *fiber.Ctx) error {
        return c.SendString("Server is UP!!")
    })

    //api.Post("/register", )

}
