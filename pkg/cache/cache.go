package cache

import (
	"time"

	"github.com/fairytale5571/secExt/pkg/cache/memory_cache"
	"github.com/fairytale5571/secExt/pkg/errs"
)

const (
	ttlMemory = 60 * time.Minute
)

type Cache interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Delete(key string)
}

type Config struct {
	memory *memory_cache.Memory
}

func SetupCache() *Config {
	return &Config{
		memory: memory_cache.New(ttlMemory),
	}
}

func (c *Config) Get(key string) (string, error) {
	if res, err := c.memory.Get(key); err == nil {
		return res, nil
	}
	return "", errs.ErrorNotCached
}

func (c *Config) Set(key, value string) error {
	if err := c.memory.Set(key, value); err != nil {
		return errs.ErrorCantCacheMemory
	}
	return nil
}
