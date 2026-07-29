package repository

import "fmt"

// ObjectRule defines event driven automation rule.
type ObjectRule struct {
	Name string `yaml:"name"`

	Event string `yaml:"event"`

	Action string `yaml:"action"`
}

// Validate validates rule.
func (r *ObjectRule) Validate() error {

	if r == nil {
		return fmt.Errorf(
			"rule is nil",
		)
	}

	if r.Name == "" {
		return fmt.Errorf(
			"rule name required",
		)
	}

	if r.Event == "" {
		return fmt.Errorf(
			"rule event required",
		)
	}

	if r.Action == "" {
		return fmt.Errorf(
			"rule action required",
		)
	}

	return nil
}
