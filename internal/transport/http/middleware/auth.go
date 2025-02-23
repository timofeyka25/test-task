package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"test-task/pkg/jwt"
)

func AuthMiddleware(accessSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Route() == nil {
			return c.Next()
		}

		tokenString := c.Get("Authorization")
		if tokenString == "" {
			log.Println("missing auth token")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing or invalid token",
			})
		}

		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		claims, err := jwt.ValidateAccessToken(tokenString, accessSecret)
		if err != nil {
			log.Println("invalid token:", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		c.Locals("user_id", claims["user_id"])

		return c.Next()
	}
}
