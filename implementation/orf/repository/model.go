package repository

import (
	"fmt"
	"strings"
)

// DefaultObjectType is used when type is not specified.
const DefaultObjectType = "object"

// DefaultObjectVersion is the initial schema version.
const DefaultObjectVersion = 1

// ObjectDefinition describes an object definition.
type ObjectDefinition struct {
	Name string `yaml:"name"`

	Type string `yaml:"type,omitempty"`

	ID string `yaml:"id,omitempty"`

	Version int `yaml:"version,omitempty"`
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

	if d.Version == 0 {
		d.Version = DefaultObjectVersion
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

	if d.Version < 1 {
		return fmt.Errorf("object version must be positive")
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
