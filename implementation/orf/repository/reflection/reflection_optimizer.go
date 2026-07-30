package reflection

type ReflectionOptimizer struct{}

func NewReflectionOptimizer() *ReflectionOptimizer {

	return &ReflectionOptimizer{}
}

func (r *ReflectionOptimizer) Optimize(
	reflection Reflection,
) *Reflection {

	if reflection.Confidence < 0.6 {

		reflection.Learning =
			"collect more data before strategy change"
	}

	reflection.Version++

	return &reflection
}
