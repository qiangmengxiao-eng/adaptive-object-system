package repository

// LifecycleRule defines object lifecycle transition.
type LifecycleRule struct {
	Name string `yaml:"name"`

	Event string `yaml:"event"`

	Action string `yaml:"action"`
}

// Validate validates lifecycle rule.
func (r LifecycleRule) Validate() error {

	if r.Name == "" {

		return ErrInvalidRule
	}

	if r.Event == "" {

		return ErrInvalidRule
	}

	if r.Action == "" {

		return ErrInvalidRule
	}

	return nil
}
