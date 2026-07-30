package decision

import "sort"

type StrategyRanker struct{}

func NewStrategyRanker() *StrategyRanker {

	return &StrategyRanker{}
}

func (s *StrategyRanker) Rank(
	options []string,
	factors map[string]float64,
) []string {

	result :=
		make([]string, len(options))

	copy(
		result,
		options,
	)

	sort.Slice(
		result,
		func(i, j int) bool {

			return score(
				result[i],
				factors,
			) > score(
				result[j],
				factors,
			)
		},
	)

	return result
}

func score(
	strategy string,
	factors map[string]float64,
) float64 {

	base :=
		0.5

	switch strategy {

	case "launch":

		base += factors["market"]

	case "optimize":

		base += factors["performance"]

	case "differentiate":

		base += factors["competition"]

	case "exit":

		base -= factors["risk"]
	}

	return base
}
