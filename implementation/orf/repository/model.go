package repository

import (
	"fmt"
	"strings"
)

// DefaultObjectType is used when type is not specified.
const DefaultObjectType = "object"

// ObjectDefinition describes an object definition.
type ObjectDefinition struct {
	Name string `yaml:"name"`

	Type string `yaml:"type,omitempty"`

	ID string `yaml:"id,omitempty"`
}

// Normalize prepares fields into canonical form.
func (d *ObjectDefinition) Normalize() {
	if d == nil {
		return
	}

	d.Name = strings.TrimSpace(d.Name)
	d.Type = strings.TrimSpace(d.Type)

	if d.Type == "" {
		d.Type = DefaultObjectType
	}

	if d.ID == "" {
		d.ID = d.Name
	}
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

// Prepare normalizes and validates the definition.
func (d *ObjectDefinition) Prepare() error {
	if d == nil {
		return fmt.Errorf("object definition is nil")
	}

	d.Normalize()

	return d.Validate()
}
