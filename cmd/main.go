package main

import (
	"fmt"
	"go-template/internal/config"
	"go-template/internal/domain/user"
	"go-template/internal/infra/database"
	"go-template/internal/infra/http"
	"go-template/internal/infra/queue"
	"go-template/pkg/shutdown"
)

func main() {
	defer shutdown.Gracefully()
	envs := config.MakeEnvs()
	db := database.MakeDatabase(envs)
	app := http.MakeServer(db)
	handler := user.QueueHandler
	queue.MakeConsumer(envs, "user", handler)
	app.Listen(fmt.Sprintf(":%s", envs.PORT))
}
