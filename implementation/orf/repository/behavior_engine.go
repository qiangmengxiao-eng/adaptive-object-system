package repository

import "fmt"

// BehaviorEngine executes object behaviors.
type BehaviorEngine struct {
	behaviors map[string]ObjectBehavior
}

// NewBehaviorEngine creates a behavior engine.
func NewBehaviorEngine() *BehaviorEngine {
	engine := &BehaviorEngine{
		behaviors: make(map[string]ObjectBehavior),
	}

	_ = engine.Register(ObjectBehavior{
		Name:   "default",
		Action: "noop",
	})

	return engine
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

// Run executes a behavior.
func (e *BehaviorEngine) Run(name string) error {
	_, ok := e.Get(name)

	if !ok {
		return fmt.Errorf("behavior not found: %s", name)
	}

	return nil
}
