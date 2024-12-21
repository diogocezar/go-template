package user

import (
	"go-template/internal/infra/queue"
	"log"
)

type Queue struct {
	producer *queue.Producer
}

func QueueHandler(body []byte) error {
	log.Printf("Processing message: %s", body)
	// Implements here!
	return nil
}
