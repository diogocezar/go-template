package user

import (
	middleware "go-template/internal/infra/http/middlaware"

	"github.com/gofiber/fiber/v2"
)

func NewRoutes(app *fiber.App, userHandler *UserHandler) {
	app.Post("/user", middleware.JWTProtected, userHandler.Create)
	app.Get("/user", userHandler.FindAll)
	app.Get("/user/:id", userHandler.FindOne)
	app.Put("/user/:id", userHandler.Update)
	app.Delete("/user/:id", userHandler.Delete)
}
