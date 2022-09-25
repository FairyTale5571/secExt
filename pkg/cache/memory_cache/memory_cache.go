package memory_cache

import (
	"time"

	"github.com/fairytale5571/secExt/pkg/errs"
	gocache "github.com/patrickmn/go-cache"
)

type Memory struct {
	client *gocache.Cache
}

func New(ttl time.Duration) *Memory {
	gocacheClient := gocache.New(ttl, ttl+15*time.Second)
	return &Memory{
		client: gocacheClient,
	}
}

func (m *Memory) Get(key string) (string, error) {
	value, found := m.client.Get(key)
	if !found {
		return "", errs.ErrorNotCached
	}
	return value.(string), nil
}

func (m *Memory) Set(key, value string) error {
	m.client.Set(key, value, gocache.DefaultExpiration)
	return nil
}

func (m *Memory) Delete(key string) {
	m.client.Delete(key)
}
