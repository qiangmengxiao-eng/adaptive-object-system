package planning

type PlanOptimizer struct{}

func NewPlanOptimizer() *PlanOptimizer {

	return &PlanOptimizer{}
}

func (p *PlanOptimizer) Optimize(
	plan Plan,
) Plan {

	if len(plan.Steps) > 3 {

		plan.Confidence += 0.05
	}

	if plan.Confidence > 1 {

		plan.Confidence = 1
	}

	return plan
}
