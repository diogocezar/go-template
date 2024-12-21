package main

import (
	"fmt"
	"go-template/internal/config"
	"go-template/internal/infra/database"
	"go-template/internal/infra/http"
	"go-template/pkg/shutdown"
)

func main() {
	defer shutdown.Gracefully()
	envs := config.MakeEnvs()
	db := database.MakeDatabase(envs)
	app := http.MakeServer(db)
	app.Listen(fmt.Sprintf(":%s", envs.PORT))
}
