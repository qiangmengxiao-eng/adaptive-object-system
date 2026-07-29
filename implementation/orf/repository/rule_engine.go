package repository

// RuleEngine evaluates events.
type RuleEngine struct {
	rules []ObjectRule
}

// NewRuleEngine creates rule engine.
func NewRuleEngine() *RuleEngine {

	return &RuleEngine{

		rules: make([]ObjectRule, 0),
	}
}

// Register adds rule.
func (e *RuleEngine) Register(
	rule ObjectRule,
) error {

	if err := rule.Validate(); err != nil {
		return err
	}

	e.rules =
		append(
			e.rules,
			rule,
		)

	return nil
}

// Match finds rules for event.
func (e *RuleEngine) Match(
	event ObjectEvent,
) []ObjectRule {

	result :=
		make([]ObjectRule, 0)

	for _, rule := range e.rules {

		if rule.Event == event.Type {

			result =
				append(
					result,
					rule,
				)
		}
	}

	return result
}

// Rules returns all rules.
func (e *RuleEngine) Rules() []ObjectRule {

	return e.rules
}
