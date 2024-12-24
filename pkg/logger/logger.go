package logger

import "github.com/rs/zerolog/log"

type Logger interface {
	Info(msg string)
	Error(msg string)
}

func Info(msg string) {
	log.Info().Msg(msg)
}

func Error(msg string) {
	log.Error().Msg(msg)
}
