package repository

// BehaviorEngine executes object behaviors.
type BehaviorEngine struct {
	behaviors map[string]ObjectBehavior
}

// NewBehaviorEngine creates a behavior engine.
func NewBehaviorEngine() *BehaviorEngine {
	return &BehaviorEngine{
		behaviors: make(map[string]ObjectBehavior),
	}
}

// Register registers a behavior.
func (e *BehaviorEngine) Register(behavior ObjectBehavior) error {
	if err := behavior.Validate(); err != nil {
		return err
	}

	e.behaviors[behavior.Name] = behavior

	return nil
}

// Get returns a behavior.
func (e *BehaviorEngine) Get(name string) (ObjectBehavior, bool) {
	behavior, ok := e.behaviors[name]

	return behavior, ok
}
