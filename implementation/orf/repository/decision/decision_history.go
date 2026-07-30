package decision

type DecisionHistory struct {
	Object string `json:"object"`

	Decisions []Decision `json:"decisions"`
}

func NewDecisionHistory(
	object string,
) *DecisionHistory {

	return &DecisionHistory{

		Object: object,

		Decisions: []Decision{},
	}
}

func (d *DecisionHistory) Add(
	decision Decision,
) {

	d.Decisions =
		append(
			d.Decisions,
			decision,
		)
}
