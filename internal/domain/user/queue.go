package user

import (
	"fmt"
	"go-template/pkg/logger"
)

func NewQueueHandler(body []byte) error {
	logger.Info(fmt.Sprintf("Processing message: %s", body))
	// Implements here!
	return nil
}
