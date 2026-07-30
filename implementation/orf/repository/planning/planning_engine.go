package planning

type PlanningEngine struct {
	Strategy *PlannerStrategy

	Optimizer *PlanOptimizer

	History *PlanHistory
}

func NewPlanningEngine() *PlanningEngine {

	return &PlanningEngine{

		Strategy: NewPlannerStrategy(),

		Optimizer: NewPlanOptimizer(),

		History: NewPlanHistory(),
	}
}

func (p *PlanningEngine) CreatePlan(
	object string,
	goal string,
) Plan {

	steps :=
		p.Strategy.Generate(
			goal,
		)

	plan :=
		Plan{

			Object: object,

			Goal: goal,

			Steps: steps,

			Strategy: "autonomous execution",

			Confidence: 0.8,

			Status: "created",

			Version: 1,

			Priority: 1,

			Risk: "medium",

			ExpectedOutcome: "achieve target goal",
		}

	optimized :=
		p.Optimizer.Optimize(
			plan,
		)

	p.History.Add(
		optimized,
	)

	return optimized
}
