package http

import (
	"go-template/internal/domain/user"
	"go-template/internal/infra/database"

	"github.com/gofiber/fiber/v2"
)

func MakeServer(database *database.Database) (app *fiber.App) {
	app = fiber.New()

	app.Use("*", func(ctx *fiber.Ctx) error {
		ctx.Set("Content-Type", "application/json")

		return ctx.Next()
	})

	userReposirory := user.MakeReposirory(database)
	userController := user.MakeController(userReposirory)
	user.MakeRoutes(app, userController)

	return app
}
