package repository

import "fmt"

// ObjectDefinition describes an object definition.
type ObjectDefinition struct {
	Name string `yaml:"name"`

	Type string `yaml:"type,omitempty"`
}

// Validate checks whether the object definition is valid.
func (d *ObjectDefinition) Validate() error {
	if d == nil {
		return fmt.Errorf("object definition is nil")
	}

	if d.Name == "" {
		return fmt.Errorf("object name is required")
	}

	return nil
}
