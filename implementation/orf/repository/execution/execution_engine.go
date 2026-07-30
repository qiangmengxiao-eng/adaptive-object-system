package execution

type ExecutionEngine struct {
	History *ExecutionHistory

	Optimizer *ExecutionOptimizer

	Feedback *ExecutionFeedback
}

func NewExecutionEngine() *ExecutionEngine {

	return &ExecutionEngine{

		History: NewExecutionHistory(),

		Optimizer: NewExecutionOptimizer(),

		Feedback: NewExecutionFeedback(),
	}
}

func (e *ExecutionEngine) Execute(
	object string,
	plan string,
	action string,
) Execution {

	execution :=
		NewExecution(
			object,
			plan,
			action,
		)

	result :=
		"action executed successfully"

	confidence :=
		0.8

	execution.Complete(
		result,
		confidence,
	)

	optimized :=
		e.Optimizer.Optimize(
			*execution,
		)

	e.History.Add(
		optimized,
	)

	return optimized
}
