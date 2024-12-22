package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shimodii/ecommerce-backend/controllers"
	"github.com/shimodii/ecommerce-backend/middleware"
)

func SetupRouter(app *fiber.App){
    // api for public apis
    api := app.Use("api")

    // admin for admin apis
    admin := app.Use("admin", middleware.AdminMiddleware)

    // admin routes
    admin.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("welcome admin")
    })
    admin.Post("/new-product", controllers.AddProduct)
    admin.Put("/update-product", controllers.UpdateProduct)
    admin.Delete("/delete-product", controllers.DeleteProduct)

    // public routes
    api.Get("/all-products", controllers.GetAllProducts)
    api.Get("/product/:id", controllers.GetProduct)

    //healthcheck
    app.Get("/status", func(c *fiber.Ctx) error {
        return c.SendString("Server is UP!!")
    })
}
