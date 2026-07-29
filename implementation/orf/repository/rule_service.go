package repository

// RuleService connects events and behaviors.
type RuleService struct {
	engine *RuleEngine

	behavior *BehaviorEngine
}

// NewRuleService creates service.
func NewRuleService(
	engine *RuleEngine,
	behavior *BehaviorEngine,
) *RuleService {

	return &RuleService{

		engine: engine,

		behavior: behavior,
	}
}

// Handle processes event.
func (s *RuleService) Handle(
	event ObjectEvent,
) []ObjectRule {

	return s.engine.Match(
		event,
	)
}
