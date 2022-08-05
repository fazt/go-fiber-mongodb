package main

import (
	"github.com/faztweb/go-fiber-mongodb/config"
	"github.com/faztweb/go-fiber-mongodb/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initializations
	app := fiber.New()
	config.ConnectDB()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"message": "pong!"})
	})

	routes.UserRoute(app)

	// Static files
	app.Static("/", "./public")

	app.Listen(":4000")

}
