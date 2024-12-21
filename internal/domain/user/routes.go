package user

import "github.com/gofiber/fiber/v2"

func MakeRoutes(app *fiber.App, controller *Controller) {
	app.Post("/user", controller.Create)
	app.Get("/user", controller.FindAll)
	app.Get("/user/:id", controller.FindOne)
	app.Put("/user", controller.Update)
	app.Delete("/user/:id", controller.Delete)
}
