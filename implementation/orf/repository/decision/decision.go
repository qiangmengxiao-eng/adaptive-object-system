package decision

type Decision struct {
	Object string `json:"object"`

	Input string `json:"input"`

	Options []string `json:"options"`

	Selected string `json:"selected"`

	Action string `json:"action"`

	Strategy string `json:"strategy"`

	Confidence float64 `json:"confidence"`

	Reason string `json:"reason"`

	Status string `json:"status"`

	Version int `json:"version"`
}

func NewDecision(
	object string,
	input string,
) *Decision {

	return &Decision{

		Object: object,

		Input: input,

		Options: []string{},

		Status: "created",

		Version: 1,
	}
}

func (d *Decision) AddOption(
	option string,
) {

	d.Options =
		append(
			d.Options,
			option,
		)
}

func (d *Decision) Select(
	strategy string,
	confidence float64,
	reason string,
) {

	d.Selected =
		strategy

	d.Strategy =
		strategy

	d.Confidence =
		confidence

	d.Reason =
		reason

	d.Status =
		"completed"
}
