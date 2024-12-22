package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shimodii/ecommerce-backend/config"
	"github.com/shimodii/ecommerce-backend/entities"
	"github.com/shimodii/ecommerce-backend/repository"
)

func LoginController(c *fiber.Ctx) error {
    var loginRequest struct {
        Email string `json:"email"`
        Password string `json:"password"`
    }

    err := c.BodyParser(&loginRequest)
    if err != nil {
        return c.JSON("invalid request")
    }
    
    var user entities.User

    err2 := repository.Database.Db.Find(&user, "email = ?", loginRequest.Email, "password = ?", loginRequest.Password)
    if err2 != nil {
        return c.Status(404).JSON("invalid crendential")
    }

    token, err := config.GenerateJwt(user)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to generate token",
        })
    }

    return c.JSON(fiber.Map{
        "token": token,
    })
}
