package server

import (
	"boilerplate/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func New() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	app.Use(func(ctx *fiber.Ctx) error {
		err := ctx.Next()
		if err != nil {
			return ctx.Status(404).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return nil
	})

	api := app.Group("/")
	routes.SetupRoutes(api)

	return app
}
