package envs

import (
	"fmt"
	"go-template/pkg/logger"
	"os"
)

func GetEnvOrDie(env string) (val string) {
	val = os.Getenv(env)

	if val == "" {
		logger.Error(fmt.Sprintf("missing environment variable %s", env))
	}

	return val
}
