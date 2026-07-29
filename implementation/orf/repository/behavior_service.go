package repository

// BehaviorService manages behavior execution.
type BehaviorService struct {
	engine *BehaviorEngine
}

// NewBehaviorService creates service.
func NewBehaviorService(
	engine *BehaviorEngine,
) *BehaviorService {

	return &BehaviorService{

		engine: engine,
	}
}

// Execute executes behavior.
func (s *BehaviorService) Execute(
	name string,
	object string,
) error {

	return s.engine.Execute(
		name,
		object,
	)
}

// Register registers behavior.
func (s *BehaviorService) Register(
	behavior ObjectBehavior,
) error {

	return s.engine.Register(
		behavior,
	)
}
