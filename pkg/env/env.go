package env

import (
	"github.com/fairytale5571/secExt/pkg/logger"
	"os"
)

type Env struct {
	logger *logger.Wrapper
}

func New() *Env {
	return &Env{
		logger: logger.New("env"),
	}
}

func (e *Env) Get(key string) string {
	if lookup, ok := os.LookupEnv(key); ok {
		return lookup
	}
	return ""
}

func (e *Env) Set(key, value string) error {
	return os.Setenv(key, value)
}
