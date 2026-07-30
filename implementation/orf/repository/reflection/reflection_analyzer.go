package reflection

type ReflectionAnalyzer struct{}

func NewReflectionAnalyzer() *ReflectionAnalyzer {

	return &ReflectionAnalyzer{}
}

func (r *ReflectionAnalyzer) Analyze(
	reflection Reflection,
) string {

	if reflection.Success {

		return "positive"

	}

	return "negative"
}
