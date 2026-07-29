package repository

import "time"

// Plan represents object execution plan.
type Plan struct {
	Object string `json:"object"`

	Intent string `json:"intent"`

	Name string `json:"name"`

	Steps []string `json:"steps"`

	Status string `json:"status"`

	Version int `json:"version"`

	CreatedAt time.Time `json:"created_at"`
}

// PlanEngine manages object plans.
type PlanEngine struct {
	plans []Plan
}

// NewPlanEngine creates plan engine.
func NewPlanEngine() *PlanEngine {

	return &PlanEngine{

		plans: make(
			[]Plan,
			0,
		),
	}
}

// Create creates a plan.
func (e *PlanEngine) Create(
	object string,
	intent string,
	name string,
	steps []string,
) Plan {

	plan :=
		Plan{

			Object: object,

			Intent: intent,

			Name: name,

			Steps: steps,

			Status: "active",

			Version: 1,

			CreatedAt: time.Now(),
		}

	e.plans =
		append(
			e.plans,
			plan,
		)

	return plan
}

// Get returns object plans.
func (e *PlanEngine) Get(
	object string,
) []Plan {

	result :=
		make(
			[]Plan,
			0,
		)

	for _, plan := range e.plans {

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
