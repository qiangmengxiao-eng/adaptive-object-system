package repository

// PolicyEngine controls object actions.
type PolicyEngine struct{}

// NewPolicyEngine creates policy engine.
func NewPolicyEngine() *PolicyEngine {

	return &PolicyEngine{}
}

// Check validates action permission.
func (p *PolicyEngine) Check(
	object string,
	action string,
) bool {

	if action == "delete" {

		return false
	}

	return true
}
