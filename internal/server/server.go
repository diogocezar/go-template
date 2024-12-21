package server

import (
	"github.com/gofiber/fiber/v2"

	"go-template/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "go-template",
			AppName:      "go-template",
		}),

		db: database.New(),
	}

	return server
}
