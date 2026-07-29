package repository

type GrowthLearning struct {
	Object string `json:"object"`

	Strategy string `json:"strategy"`

	Confidence float64 `json:"confidence"`

	Recommendation string `json:"recommendation"`
}

type GrowthLearningEngine struct {
	Reflection *PerformanceReflectionEngine
}

func NewGrowthLearningEngine(
	r *PerformanceReflectionEngine,
) *GrowthLearningEngine {

	return &GrowthLearningEngine{

		Reflection: r,
	}
}

func (g *GrowthLearningEngine) Learn(
	performance SalesPerformance,
) GrowthLearning {

	reflection :=
		g.Reflection.Reflect(
			performance,
		)

	confidence :=
		0.5

	if reflection.Problem ==
		"performance acceptable" {

		confidence =
			0.9
	}

	return GrowthLearning{

		Object: performance.Object,

		Strategy: reflection.Action,

		Confidence: confidence,

		Recommendation: reflection.Reason,
	}
}
