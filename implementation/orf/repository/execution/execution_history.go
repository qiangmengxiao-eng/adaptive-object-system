package execution

type ExecutionHistory struct {
	Executions []Execution `json:"executions"`
}

func NewExecutionHistory() *ExecutionHistory {

	return &ExecutionHistory{

		Executions: []Execution{},
	}
}

func (h *ExecutionHistory) Add(
	execution Execution,
) {

	h.Executions =
		append(
			h.Executions,
			execution,
		)
}

func (h *ExecutionHistory) Latest() *Execution {

	if len(h.Executions) == 0 {

		return nil
	}

	return &h.Executions[len(h.Executions)-1]
}

func (h *ExecutionHistory) Count() int {

	return len(
		h.Executions,
	)
}
