package repository

import "fmt"

// ObjectEvent describes an object system event.
type ObjectEvent struct {
	Type string `yaml:"type"`

	Object string `yaml:"object"`

	Data string `yaml:"data,omitempty"`
}

// Validate validates event.
func (e *ObjectEvent) Validate() error {

	if e == nil {
		return fmt.Errorf("event is nil")
	}

	if e.Type == "" {
		return fmt.Errorf("event type required")
	}

	if e.Object == "" {
		return fmt.Errorf("event object required")
	}

	return nil
}
