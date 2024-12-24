package health

import "github.com/gofiber/fiber/v2"

type HealthHandler struct{}

func NewHandler() *HealthHandler {
	return &HealthHandler{}
}

type HealthCheckResponse struct {
	Status string `json:"status"`
}

// Liveness Health Check godoc
//
//	@Summary		Check if application is alive
//	@Description	Check if application is alive
//	@Tags			health-check
//	@Success		200	{object}	HealthCheckResponse	"Application is alive"
//	@Failure		500	{object}	interface{}	"Internal Server Error"
//	@Router			/healthz [get]
func (handler *HealthHandler) Liveness(ctx *fiber.Ctx) error {
	return ctx.
		Status(fiber.StatusOK).
		JSON(HealthCheckResponse{Status: "alive"})
}

// Readiness Health Check godoc
//
//	@Summary		Check if application is ready
//	@Description	Check if application is ready
//	@Tags			health-check
//	@Success		200	{object}	HealthCheckResponse	"Application is ready"
//	@Failure		503	{object}	HealthCheckResponse	"Service Unavailable"
//	@Router			/healthz/ready [get]
func (handler *HealthHandler) Readiness(ctx *fiber.Ctx) error {
	if isValid := dummyValidateServices(); !isValid {
		return ctx.
			Status(fiber.StatusServiceUnavailable).
			JSON(HealthCheckResponse{Status: "unavailable"})
	}

	return ctx.
		Status(fiber.StatusOK).
		JSON(HealthCheckResponse{Status: "ready"})
}

// dummy function, to simulate services health
func dummyValidateServices() bool {
	return false
}
