package http

import (
	"go-template/internal/domain/user"
	"go-template/internal/infra/database"
	"go-template/internal/infra/queue"

	"github.com/gofiber/fiber/v2"
)

func MakeServer(database *database.Database, producer *queue.Producer) (app *fiber.App) {
	app = fiber.New()

	app.Use("*", func(ctx *fiber.Ctx) error {
		ctx.Set("Content-Type", "application/json")

		return ctx.Next()
	})

	userReposirory := user.MakeReposirory(database)
	userController := user.MakeController(userReposirory, producer)
	user.MakeRoutes(app, userController)

	return app
}
