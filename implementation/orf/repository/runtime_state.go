package repository

import "fmt"

// RuntimeState represents object lifecycle state.
type RuntimeState struct {
	Name string `yaml:"name"`

	Status string `yaml:"status"`
}

// NewRuntimeState creates state.
func NewRuntimeState(
	name string,
	status string,
) RuntimeState {

	return RuntimeState{
		Name:   name,
		Status: status,
	}
}

// Validate validates state.
func (s *RuntimeState) Validate() error {

	if s == nil {
		return fmt.Errorf(
			"runtime state is nil",
		)
	}

	if s.Name == "" {
		return fmt.Errorf(
			"state name required",
		)
	}

	if s.Status == "" {
		return fmt.Errorf(
			"state status required",
		)
	}

	return nil
}
