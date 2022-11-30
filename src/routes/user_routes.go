package routes

import (
	"github.com/faztweb/go-fiber-mongodb/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/users", controllers.CreateUser)
	app.Get("/users/:userId", controllers.GetUser)
	app.Delete("/users/:userId", controllers.DeleteUser)
	app.Get("/users", controllers.GetUsers)
	app.Patch("/users/:userId", controllers.UpdateUser)
}
