package main

import (
	"fmt"
	"go-template/internal/config"
	"go-template/internal/domain/user"
	"go-template/internal/infra/database"
	"go-template/internal/infra/http"
	"go-template/internal/infra/queue"
	"go-template/pkg/shutdown"
	"log"
)

func main() {
	defer shutdown.Gracefully()
	envs := config.MakeEnvs()
	db := database.MakeDatabase(envs)
	producer, err := queue.MakeProducer(envs)
	if err != nil {
		log.Fatalf("Error trying to connect on Queue: %v", err)
	}
	defer queue.CloseProducer(producer)
	handler := user.QueueHandler
	go func() {
		queue.MakeConsumer(envs, "users", handler)
	}()
	app := http.MakeServer(db, producer)
	app.Listen(fmt.Sprintf(":%s", envs.PORT))
}
