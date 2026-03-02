package main

import (
	"back-end-hr-manajement/internal/config"
	"back-end-hr-manajement/internal/connection"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	config := config.Get()
	database := connection.GetDatabaseConnection(config.Database)
	_ = database
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "HRIS Backend Running 🚀",
		})
	})

	log.Fatal(app.Listen(":9000"))
}
