package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shimodii/ecommerce-backend/middleware"
)

func SetupRoutes(app *fiber.App){
    api := app.Group("api")
    admin := api.Group("admin",middleware.AdminMiddleware)
    app.Get("/status",func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "message": "success",
        })
    })       
    admin.Post("/new",cotr) 


}
