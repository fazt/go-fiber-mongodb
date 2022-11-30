package main

import (
	"github.com/faztweb/go-fiber-mongodb/src/config"
	"github.com/faztweb/go-fiber-mongodb/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initializations
	app := fiber.New()
	config.ConnectDB()

	// Routes
	routes.IndexRoutes(app)
	routes.UserRoute(app)

	// Static files
	app.Static("/", "./public")

	app.Listen(":3000")
}
