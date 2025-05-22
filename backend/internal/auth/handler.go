package auth

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input User
		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
		}

		hash, err := HashPassword(input.Password)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Error hashing password"})
		}

		user := User{Name: input.Name, Email: input.Email, Password: hash}
		if err := db.Create(&user).Error; err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Email already exists"})
		}

		token, _ := GenerateJWT(user.ID)
		return c.JSON(fiber.Map{"token": token})
	}
}

func Login(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input User
		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
		}

		var user User
		if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid email or password"})
		}

		if !CheckPasswordHash(input.Password, user.Password) {
			return c.Status(401).JSON(fiber.Map{"error": "Invalid email or password"})
		}

		token, _ := GenerateJWT(user.ID)
		return c.JSON(fiber.Map{"token": token})
	}
}

func Me(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(User)
		return c.JSON(user)
	}
}
