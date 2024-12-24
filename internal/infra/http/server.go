package http

import (
	"go-template/internal/domain/auth"
	"go-template/internal/domain/health"
	"go-template/internal/domain/user"
	"go-template/internal/infra/database"
	"go-template/internal/infra/queue"

	_ "go-template/api/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func New(database *database.Database, producer *queue.Producer) (app *fiber.App) {
	app = fiber.New()

	app.Use("*", func(ctx *fiber.Ctx) error {
		ctx.Set("Content-Type", "application/json")

		return ctx.Next()
	})

	app.Get("/swagger/*", swagger.HandlerDefault)

	userReposirory := user.NewRepository(database)
	userHandler := user.NewHandler(userReposirory, producer)
	user.NewRoutes(app, userHandler)

	authHandler := auth.NewHandler(userReposirory)
	auth.NewRoutes(app, authHandler)

	healthHandler := health.NewHandler()
	health.NewRoutes(app, healthHandler)

	return app
}
