package repository

import "time"

// GeneratedPlan represents an automatically generated plan.
type GeneratedPlan struct {
	Object string `json:"object"`

	Decision string `json:"decision"`

	Name string `json:"name"`

	Steps []string `json:"steps"`

	Status string `json:"status"`

	Version int `json:"version"`

	CreatedAt time.Time `json:"created_at"`
}

// PlannerEngine creates plans from decisions.
type PlannerEngine struct {
	plans []GeneratedPlan
}

// NewPlannerEngine creates planner engine.
func NewPlannerEngine() *PlannerEngine {

	return &PlannerEngine{

		plans: make(
			[]GeneratedPlan,
			0,
		),
	}
}

// Generate creates a plan according to decision.
func (p *PlannerEngine) Generate(
	object string,
	decision Decision,
) GeneratedPlan {

	name :=
		decision.Strategy

	steps :=
		[]string{
			"observe current state",
			"execute next action",
			"evaluate result",
		}

	switch decision.Strategy {

	case "retry_strategy":

		name =
			"retry_strategy"

		steps =
			[]string{
				"analyze failure",
				"retry task",
				"record result",
			}

	case "adaptation_strategy":

		name =
			"adaptation_strategy"

		steps =
			[]string{
				"analyze failure",
				"adjust behavior",
				"retry task",
			}

	case "observe_strategy":

		name =
			"observe_strategy"

		steps =
			[]string{
				"observe environment",
				"collect experience",
				"evaluate state",
			}

	case "continue_strategy":

		name =
			"continue_strategy"

	default:

		name =
			"continue_strategy"
	}

	plan :=
		GeneratedPlan{

			Object: object,

			Decision: decision.Type,

			Name: name,

			Steps: steps,

			Status: "active",

			Version: 1,

			CreatedAt: time.Now(),
		}

	p.plans =
		append(
			p.plans,
			plan,
		)

	return plan
}

// Get returns generated plans.
func (p *PlannerEngine) Get(
	object string,
) []GeneratedPlan {

	result :=
		make(
			[]GeneratedPlan,
			0,
		)

	for _, plan := range p.plans {

		if plan.Object ==
			object {

			result =
				append(
					result,
					plan,
				)
		}
	}

	return result
}
