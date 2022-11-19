package uuid

import (
	guuid "github.com/google/uuid"
)

// GenerateUUID generate random uuid
func GenerateUUID() string {
	return guuid.New().String()
}

// Provider provide uuid dependency.
type Provider interface {
	Generate() string
}

// UUID ...
type UUID struct{}

// New ...
func New() *UUID {
	return new(UUID)
}

// Generate generate random uuid
func (*UUID) Generate() string {
	return guuid.New().String()
}
