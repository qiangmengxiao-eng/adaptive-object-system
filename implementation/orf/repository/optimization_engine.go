package repository

type OptimizationEngine struct {
	Decision *DecisionEngine
}

func NewOptimizationEngine(
	d *DecisionEngine,
) *OptimizationEngine {

	return &OptimizationEngine{
		Decision: d,
	}
}

func (o *OptimizationEngine) Optimize(
	object string,
) string {

	decision, _ :=
		o.Decision.Decide(
			object,
		)

	if decision.Type == "continue" {

		return "stable"
	}

	return "adapt"
}
