package main

import (
	"log"
	"os"

	"github.com/AAbregoAndazola/Go-tasker/internal/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()

	dsn := os.Getenv("DB_DSN") // ejemplo: postgres://user:pass@localhost:5432/tasker?sslmode=disable
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&auth.User{})

	app := fiber.New()

	api := app.Group("/api")

	api.Post("/register", auth.Register(db))
	api.Post("/login", auth.Login(db))
	api.Get("/me", auth.Protected(db), auth.Me(db))

	app.Listen(":3000")
}
