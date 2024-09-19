package server

import "github.com/gofiber/fiber/v2"

func initializeRoutes(app *fiber.App) *fiber.App {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	return app
}
