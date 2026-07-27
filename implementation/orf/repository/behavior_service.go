package repository

// BehaviorService provides behavior operations.
type BehaviorService struct {
	engine *BehaviorEngine
}

// NewBehaviorService creates behavior service.
func NewBehaviorService(engine *BehaviorEngine) *BehaviorService {
	return &BehaviorService{
		engine: engine,
	}
}

// List returns available behaviors.
func (s *BehaviorService) List() []string {
	result := make([]string, 0)

	for name := range s.engine.behaviors {
		result = append(result, name)
	}

	return result
}

// Run executes behavior.
func (s *BehaviorService) Run(name string, input any) error {
	return s.engine.Run(name)
}
