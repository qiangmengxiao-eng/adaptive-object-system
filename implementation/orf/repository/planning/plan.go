package planning

import (
	"time"
)

type Plan struct {
	Object string `json:"object"`

	Goal string `json:"goal"`

	Steps []string `json:"steps"`

	Strategy string `json:"strategy"`

	Confidence float64 `json:"confidence"`

	Status string `json:"status"`

	Version int `json:"version"`

	// Phase 12-15 Autonomous Planning

	Priority int `json:"priority"`

	Risk string `json:"risk"`

	ExpectedOutcome string `json:"expected_outcome"`

	CreatedAt time.Time `json:"created_at"`
}

func NewPlan(
	object string,
	goal string,
) *Plan {

	return &Plan{

		Object: object,

		Goal: goal,

		Steps: []string{},

		Status: "created",

		Version: 1,

		Priority: 1,

		Risk: "unknown",

		ExpectedOutcome: "",

		CreatedAt: time.Now(),
	}
}

func (p *Plan) AddStep(
	step string,
) {

	p.Steps =
		append(
			p.Steps,
			step,
		)
}

func (p *Plan) Complete(
	strategy string,
	confidence float64,
) {

	p.Strategy =
		strategy

	p.Confidence =
		confidence

	p.Status =
		"completed"
}

func (p *Plan) SetPriority(
	priority int,
) {

	p.Priority =
		priority
}

func (p *Plan) SetRisk(
	risk string,
) {

	p.Risk =
		risk
}

func (p *Plan) SetOutcome(
	outcome string,
) {

	p.ExpectedOutcome =
		outcome
}
