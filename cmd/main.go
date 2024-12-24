package main

import (
	"fmt"
	"go-template/internal/config"
	"go-template/internal/domain/user"
	"go-template/internal/infra/database"
	"go-template/internal/infra/http"
	"go-template/internal/infra/queue"
	"go-template/pkg/logger"
	"go-template/pkg/shutdown"
)

func main() {
	logger.Info("Starting Application")
	defer shutdown.Gracefully()

	logger.Info("Getting Envs")
	envs := config.New()

	logger.Info("Creating Database Connection")
	db := database.New(envs)

	logger.Info("Creating Producer to Queue")
	producer, err := queue.NewProducer(envs)
	if err != nil {
		logger.Error(fmt.Sprintf("Error trying to connect on Queue: %v", err))
	}
	defer queue.CloseProducer(producer)

	logger.Info("Creating Handler to Consume Queue")
	queueHandler := user.NewQueueHandler
	go func() {
		queue.NewConsumer(envs, "users", queueHandler)
	}()

	logger.Info("Creating HTTP Server")
	app := http.New(db, producer)
	app.Listen(fmt.Sprintf(":%s", envs.PORT))
}
