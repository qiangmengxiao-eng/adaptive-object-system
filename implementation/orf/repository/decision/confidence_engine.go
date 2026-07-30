package decision

type ConfidenceEngine struct{}

func NewConfidenceEngine() *ConfidenceEngine {

	return &ConfidenceEngine{}
}

func (c *ConfidenceEngine) Calculate(
	factors map[string]float64,
) float64 {

	total :=
		0.0

	count :=
		0.0

	for _, value := range factors {

		total += value

		count++
	}

	if count == 0 {

		return 0
	}

	result :=
		total / count

	if result > 1 {

		result = 1
	}

	return result
}
