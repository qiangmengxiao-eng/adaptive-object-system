package repository

// Decision represents system decision.
type Decision struct {
	Object string `json:"object"`

	Type string `json:"type"`

	Reason string `json:"reason"`

	Confidence float64 `json:"confidence"`

	Strategy string `json:"strategy"`
}

// DecisionEngine evaluates experience.
type DecisionEngine struct {
	Experience *ExperienceEngine

	Reflection *ReflectionEngine

	Knowledge *KnowledgeEngine
}

// NewDecisionEngine creates decision engine.
func NewDecisionEngine(
	experience *ExperienceEngine,
) *DecisionEngine {

	return &DecisionEngine{

		Experience: experience,
	}
}

// AttachReflection attaches reflection feedback.
func (d *DecisionEngine) AttachReflection(
	reflection *ReflectionEngine,
) {

	d.Reflection =
		reflection
}

// AttachKnowledge attaches knowledge engine.
func (d *DecisionEngine) AttachKnowledge(
	knowledge *KnowledgeEngine,
) {

	d.Knowledge =
		knowledge
}

// Decide generates decision.
func (d *DecisionEngine) Decide(
	object string,
) (*Decision, error) {

	experience :=
		d.Experience.Analyze(
			object,
		)

	decision :=
		&Decision{

			Object: object,
		}

	// Knowledge has priority.
	if d.Knowledge != nil {

		knowledge :=
			d.Knowledge.Get(
				object,
			)

		for _, item := range knowledge {

			if item.Preferred {

				decision.Strategy =
					item.Strategy

				break
			}
		}
	}

	if experience.Events == 0 {

		decision.Type =
			"observe"

		decision.Reason =
			"insufficient experience"

		decision.Confidence =
			0.1

		if decision.Strategy == "" {

			decision.Strategy =
				"observe_strategy"
		}

		return decision, nil
	}

	// Reflection feedback.
	if d.Reflection != nil {

		reflection :=
			d.Reflection.Latest(
				object,
			)

		if reflection != nil &&
			reflection.SuccessRate < 0.5 {

			decision.Type =
				"adapt"

			decision.Reason =
				"reflection_detected_low_performance"

			decision.Confidence =
				0.9

			if decision.Strategy == "" {

				decision.Strategy =
					"retry_strategy"
			}

			return decision, nil
		}
	}

	// Experience based adaptation.
	if experience.SuccessRate() < 0.5 {

		decision.Type =
			"adapt"

		decision.Reason =
			"failure_rate_high"

		decision.Confidence =
			0.8

		if decision.Strategy == "" {

			decision.Strategy =
				"retry_strategy"
		}

		return decision, nil
	}

	// Continue.
	decision.Type =
		"continue"

	decision.Reason =
		"performance acceptable"

	decision.Confidence =
		0.7

	if decision.Strategy == "" {

		decision.Strategy =
			"continue_strategy"
	}

	return decision, nil
}
