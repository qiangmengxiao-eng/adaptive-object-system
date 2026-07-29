package repository

// LifecycleEngine manages lifecycle rules.
type LifecycleEngine struct {
	rules []LifecycleRule

	store *LifecycleRuleStore
}

// NewLifecycleEngine creates lifecycle engine.
func NewLifecycleEngine(
	store *LifecycleRuleStore,
) *LifecycleEngine {

	engine :=
		&LifecycleEngine{

			rules: make(
				[]LifecycleRule,
				0,
			),

			store: store,
		}

	engine.Load()

	return engine
}

// Register adds lifecycle rule.
func (e *LifecycleEngine) Register(
	rule LifecycleRule,
) error {

	if err := rule.Validate(); err != nil {

		return err
	}

	for _, existing := range e.rules {

		if existing.Name == rule.Name &&
			existing.Event == rule.Event &&
			existing.Action == rule.Action {

			return nil
		}
	}

	e.rules =
		append(
			e.rules,
			rule,
		)

	if e.store != nil {

		_ =
			e.store.Save(
				e.rules,
			)
	}

	return nil
}

// Handle processes event.
func (e *LifecycleEngine) Handle(
	event ObjectEvent,
	object *RuntimeObject,
) {

	if object == nil {

		return
	}

	for _, rule := range e.rules {

		if rule.Event != event.Type {

			continue
		}

		switch rule.Action {

		case "state.completed":

			StateTransition{

				From: object.State.Status,

				To: "completed",
			}.Apply(
				object,
			)

		case "state.active":

			StateTransition{

				From: object.State.Status,

				To: "active",
			}.Apply(
				object,
			)
		}
	}
}

// Load restores rules.
func (e *LifecycleEngine) Load() {

	if e.store == nil {

		return
	}

	rules, err :=
		e.store.Load()

	if err != nil {

		return
	}

	e.rules =
		rules
}

// List returns rules.
func (e *LifecycleEngine) List() []LifecycleRule {

	return e.rules
}
