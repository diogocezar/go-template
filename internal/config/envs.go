package config

import (
	"fmt"
	"go-template/pkg/envs"
	"go-template/pkg/logger"

	"github.com/joho/godotenv"
)

type Envs struct {
	DATABASE_HOST     string
	DATABASE_NAME     string
	DATABASE_PASSWORD string
	DATABASE_PORT     string
	DATABASE_USER     string
	PORT              string
	QUEUE_HOST        string
	QUEUE_PORT        string
	QUEUE_USER        string
	QUEUE_PASSWORD    string
	JWT_SECRET        string
	ENVIRONMENT       string
}

func New() *Envs {
	if err := godotenv.Load(".env"); err != nil {
		logger.Error(fmt.Sprintf("Error loading .env file: %v", err))
	}

	return &Envs{
		DATABASE_HOST:     envs.GetEnvOrDie("DATABASE_HOST"),
		DATABASE_NAME:     envs.GetEnvOrDie("DATABASE_NAME"),
		DATABASE_PASSWORD: envs.GetEnvOrDie("DATABASE_PASSWORD"),
		DATABASE_PORT:     envs.GetEnvOrDie("DATABASE_PORT"),
		DATABASE_USER:     envs.GetEnvOrDie("DATABASE_USER"),
		PORT:              envs.GetEnvOrDie("PORT"),
		QUEUE_HOST:        envs.GetEnvOrDie("QUEUE_HOST"),
		QUEUE_PORT:        envs.GetEnvOrDie("QUEUE_PORT"),
		QUEUE_USER:        envs.GetEnvOrDie("QUEUE_USER"),
		QUEUE_PASSWORD:    envs.GetEnvOrDie("QUEUE_PASSWORD"),
		JWT_SECRET:        envs.GetEnvOrDie("JWT_SECRET"),
		ENVIRONMENT:       envs.GetEnvOrDie("ENVIRONMENT"),
	}
}
