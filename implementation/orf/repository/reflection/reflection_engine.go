package reflection

type ReflectionEngine struct {
	Analyzer *ReflectionAnalyzer

	Optimizer *ReflectionOptimizer
}

func NewReflectionEngine() *ReflectionEngine {

	return &ReflectionEngine{

		Analyzer: NewReflectionAnalyzer(),

		Optimizer: NewReflectionOptimizer(),
	}
}

func (r *ReflectionEngine) Reflect(
	object string,
	execution string,
	result string,
) Reflection {

	reflection :=
		NewReflection(
			object,
			execution,
		)

	if result == "success" {

		reflection.Success =
			true

		reflection.Reason =
			"execution achieved target"

		reflection.Learning =
			"continue current strategy"

		reflection.Confidence =
			0.9

	} else {

		reflection.Success =
			false

		reflection.Problem =
			"execution failed"

		reflection.Reason =
			"strategy requires adjustment"

		reflection.Learning =
			"optimize next execution"

		reflection.Confidence =
			0.5

	}

	return *r.Optimizer.Optimize(
		*reflection,
	)
}
