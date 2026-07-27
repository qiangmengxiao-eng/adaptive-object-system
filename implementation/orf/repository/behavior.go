package repository

import "fmt"

// ObjectBehavior describes object behavior.
type ObjectBehavior struct {
	Name string `yaml:"name"`

	Action string `yaml:"action"`
}

// Validate checks behavior validity.
func (b *ObjectBehavior) Validate() error {
	if b == nil {
		return fmt.Errorf("object behavior is nil")
	}

	if b.Name == "" {
		return fmt.Errorf("behavior name is required")
	}

	if b.Action == "" {
		return fmt.Errorf("behavior action is required")
	}

	return nil
}
