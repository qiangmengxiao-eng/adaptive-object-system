package repository

import "fmt"

// ObjectMetadata describes object metadata.
type ObjectMetadata struct {
	Version int `yaml:"version"`

	CreatedAt string `yaml:"created_at,omitempty"`
}

// Validate checks metadata validity.
func (m *ObjectMetadata) Validate() error {
	if m == nil {
		return fmt.Errorf("object metadata is nil")
	}

	if m.Version < 1 {
		return fmt.Errorf("metadata version must be greater than zero")
	}

	return nil
}
