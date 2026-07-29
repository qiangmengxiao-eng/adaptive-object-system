package repository

// StateTransition represents runtime state change.
type StateTransition struct {
	From string `yaml:"from"`

	To string `yaml:"to"`
}

// Apply applies transition.
func (t StateTransition) Apply(
	object *RuntimeObject,
) {

	if object == nil {

		return
	}

	object.State.Status =
		t.To
}
