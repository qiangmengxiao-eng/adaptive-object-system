package repository

type LifecycleManager struct {
	Experience *ExperienceEngine
}

func NewLifecycleManager(
	e *ExperienceEngine,
) *LifecycleManager {

	return &LifecycleManager{
		Experience: e,
	}
}

func (l *LifecycleManager) Evaluate(
	object string,
) string {

	exp :=
		l.Experience.Analyze(
			object,
		)

	if exp.SuccessRate() > 0.9 {

		return "optimized"
	}

	if exp.Events > 0 {

		return "learning"
	}

	return "created"
}
