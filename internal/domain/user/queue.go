package user

import (
	"log"
)

func QueueHandler(body []byte) error {
	log.Printf("Processing message: %s", body)
	// Implements here!
	return nil
}
