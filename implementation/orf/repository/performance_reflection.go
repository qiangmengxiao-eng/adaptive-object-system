package repository

type PerformanceReflection struct {
	Problem string `json:"problem"`

	Action string `json:"action"`

	Reason string `json:"reason"`
}

type PerformanceReflectionEngine struct{}

func NewPerformanceReflectionEngine() *PerformanceReflectionEngine {

	return &PerformanceReflectionEngine{}
}

func (p *PerformanceReflectionEngine) Reflect(
	performance SalesPerformance,
) PerformanceReflection {

	if performance.CTR < 2 {

		return PerformanceReflection{

			Problem: "low click through rate",

			Action: "improve title and main image",

			Reason: "customers are not clicking",
		}
	}

	if performance.ConversionRate < 5 {

		return PerformanceReflection{

			Problem: "low conversion",

			Action: "improve listing benefits",

			Reason: "traffic does not convert",
		}
	}

	return PerformanceReflection{

		Problem: "performance acceptable",

		Action: "continue current strategy",

		Reason: "metrics within range",
	}
}
