package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Protected(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth == "" {
			return c.Status(401).JSON(fiber.Map{"error": "Missing token"})
		}

		token := strings.TrimPrefix(auth, "Bearer ")
		claims, err := ParseJWT(token)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid token"})
		}

		var user User
		if err := db.First(&user, claims.UserID).Error; err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "User not found"})
		}

		c.Locals("user", user)
		return c.Next()
	}
}
