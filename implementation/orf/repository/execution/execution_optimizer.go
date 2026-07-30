package execution

type ExecutionOptimizer struct{}

func NewExecutionOptimizer() *ExecutionOptimizer {

	return &ExecutionOptimizer{}
}

func (e *ExecutionOptimizer) Optimize(
	execution Execution,
) Execution {

	if execution.Confidence < 0.5 {

		execution.Status =
			"needs_review"
	}

	execution.Version++

	return execution
}
