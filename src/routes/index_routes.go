package routes

import "github.com/gofiber/fiber/v2"

func IndexRoutes(app *fiber.App) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"message": "pong!"})
	})
}
