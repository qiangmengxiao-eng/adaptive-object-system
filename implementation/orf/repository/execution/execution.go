package execution

type Execution struct {
	Object string `json:"object"`

	Plan string `json:"plan"`

	Action string `json:"action"`

	Result string `json:"result"`

	Status string `json:"status"`

	Confidence float64 `json:"confidence"`

	Version int `json:"version"`
}

func NewExecution(
	object string,
	plan string,
	action string,
) *Execution {

	return &Execution{

		Object: object,

		Plan: plan,

		Action: action,

		Status: "created",

		Confidence: 0.5,

		Version: 1,
	}
}

func (e *Execution) Complete(
	result string,
	confidence float64,
) {

	e.Result = result

	e.Confidence = confidence

	e.Status = "completed"
}
