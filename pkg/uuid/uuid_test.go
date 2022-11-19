package uuid

import (
	"testing"
)

type mockUUID struct{}

func (*mockUUID) Generate() string {
	return "MockID"
}

func TestGenerate(t *testing.T) {
	m := New()
	if uuid := m.Generate(); uuid == "" {
		t.Errorf("Expected non empty uuid but found uuid empty.")
	}
}
