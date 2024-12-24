package auth

import "github.com/gofiber/fiber/v2"

func NewRoutes(app *fiber.App, authHandler *AuthHandler) {
	app.Post("/auth/login", authHandler.Login)
}
