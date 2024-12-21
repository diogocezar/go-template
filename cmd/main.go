package main

import (
	"fmt"
	"go-template/internal/config"
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
	app := http.MakeServer(db, producer)
	//handler := user.QueueHandler
	//queue.MakeConsumer(envs, "users", handler)
	app.Listen(fmt.Sprintf(":%s", envs.PORT))
}
