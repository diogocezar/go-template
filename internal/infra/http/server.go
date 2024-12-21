package http

import (
	"go-template/internal/infra/database"

	"github.com/gofiber/fiber/v2"
)

func MakeServer(database *database.Database) (app *fiber.App) {
	app = fiber.New()

	app.Use("*", func(ctx *fiber.Ctx) error {
		ctx.Set("Content-Type", "application/json")

		return ctx.Next()
	})

	// peopleReposirory := people.MakeReposirory(database)
	// peopleController := people.MakeController(peopleReposirory)
	// people.MakeRoutes(app, peopleController)

	return app
}
