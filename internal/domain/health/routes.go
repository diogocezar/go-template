package health

import (
	"github.com/gofiber/fiber/v2"
)

func NewRoutes(app *fiber.App, healthHandler *HealthHandler) {
	app.Get("/healthz", healthHandler.Liveness)
	app.Get("/healthz/ready", healthHandler.Readiness)
}
