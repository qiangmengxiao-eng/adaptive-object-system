package repository

// BehaviorEngine executes object behaviors.
type BehaviorEngine struct {
	behaviors map[string]ObjectBehavior
}

// NewBehaviorEngine creates behavior engine.
func NewBehaviorEngine() *BehaviorEngine {

	return &BehaviorEngine{

		behaviors: make(
			map[string]ObjectBehavior,
		),
	}
}

// Register registers behavior.
func (e *BehaviorEngine) Register(
	behavior ObjectBehavior,
) error {

	if err :=
		behavior.Validate(); err != nil {

		return err
	}

	e.behaviors[behavior.Name] = behavior

	return nil
}

// Get returns behavior.
func (e *BehaviorEngine) Get(
	name string,
) (
	ObjectBehavior,
	bool,
) {

	behavior, ok :=
		e.behaviors[name]

	return behavior, ok
}

// Execute executes behavior.
func (e *BehaviorEngine) Execute(
	name string,
	object string,
) error {

	behavior, ok :=
		e.Get(
			name,
		)

	if !ok {

		return nil
	}

	_ = behavior

	_ = object

	return nil
}

// List returns behaviors.
func (e *BehaviorEngine) List() []ObjectBehavior {

	result :=
		make(
			[]ObjectBehavior,
			0,
		)

	for _, behavior := range e.behaviors {

		result =
			append(
				result,
				behavior,
			)
	}

	return result
}
