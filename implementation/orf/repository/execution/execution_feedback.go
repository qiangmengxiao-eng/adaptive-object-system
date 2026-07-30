package execution

type Feedback struct {
	Execution string `json:"execution"`

	Success bool `json:"success"`

	Message string `json:"message"`

	Confidence float64 `json:"confidence"`
}

type ExecutionFeedback struct{}

func NewExecutionFeedback() *ExecutionFeedback {

	return &ExecutionFeedback{}
}

func (e *ExecutionFeedback) Generate(
	execution Execution,
) Feedback {

	success :=
		execution.Status == "completed"

	return Feedback{

		Execution: execution.Action,

		Success: success,

		Message: execution.Result,

		Confidence: execution.Confidence,
	}
}
