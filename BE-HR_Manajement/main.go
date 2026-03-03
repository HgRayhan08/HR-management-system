package main

import (
	"back-end-hr-manajement/internal/api"
	"back-end-hr-manajement/internal/config"
	"back-end-hr-manajement/internal/connection"
	"back-end-hr-manajement/internal/middleware"
	"back-end-hr-manajement/internal/repository"
	"back-end-hr-manajement/internal/service"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	config := config.Get()
	database := connection.GetDatabaseConnection(config.Database)
	app := fiber.New()

	middlewareJWT := middleware.JWTProtected(*config)

	userRepository := repository.NewUserRepository(database)
	authService := service.NewAuthService(userRepository)
	api.NewAuthApi(app, authService, middlewareJWT)

	// app.Get("/", func(c fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{
	// 		"status":  "ok",
	// 		"message": "HRIS Backend Running 🚀",
	// 	})
	// })

	log.Fatal(app.Listen(":9000"))
}
