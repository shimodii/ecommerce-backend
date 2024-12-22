package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("mySecret")

func AdminMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")[7:] // Remove "Bearer "

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
		}
		return jwtKey, nil
	})

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token claims")
	}

	role, ok := claims["role"].(string)
	if !ok || role != "admin" {
		return fiber.NewError(fiber.StatusForbidden, "Access denied")
	}

	return c.Next()
}
