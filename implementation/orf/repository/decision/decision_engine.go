package decision

type DecisionEngine struct {
	Confidence *ConfidenceEngine

	Ranker *StrategyRanker
}

func NewDecisionEngine() *DecisionEngine {

	return &DecisionEngine{

		Confidence: NewConfidenceEngine(),

		Ranker: NewStrategyRanker(),
	}
}

func (d *DecisionEngine) Decide(
	object string,
	input string,
	options []string,
	factors map[string]float64,
) Decision {

	ranked :=
		d.Ranker.Rank(
			options,
			factors,
		)

	selected :=
		"unknown"

	if len(ranked) > 0 {

		selected =
			ranked[0]
	}

	confidence :=
		d.Confidence.Calculate(
			factors,
		)

	return Decision{

		Object: object,

		Input: input,

		Options: options,

		Selected: selected,

		Strategy: selected,

		Confidence: confidence,

		Reason: "selected highest ranked strategy",

		Status: "completed",

		Version: 1,
	}
}
