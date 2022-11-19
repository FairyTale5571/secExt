package mock

// UUID is a mock implementation of uuid.
type UUID struct {
	id string
}

// Generate generates the id for the uuid.
func (m *UUID) Generate() string {
	return m.id
}

// NewUUID returns mock instance of UUID
// Use it for testing only.
func NewUUID(id string) *UUID {
	return &UUID{
		id: id,
	}
}
