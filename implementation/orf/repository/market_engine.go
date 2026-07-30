package repository

type MarketAnalysisResult struct {
	Opportunity string `json:"opportunity"`

	Competition string `json:"competition"`

	Score float64 `json:"score"`

	Recommendation string `json:"recommendation"`
}

type MarketEngine struct{}

func NewMarketEngine() *MarketEngine {

	return &MarketEngine{}
}

func (m *MarketEngine) Analyze(
	searchVolume int,
	competition int,
) MarketAnalysisResult {

	score :=

		float64(searchVolume) /
			float64(competition+1)

	result :=
		MarketAnalysisResult{

			Score: score,
		}

	if score > 1000 {

		result.Opportunity =
			"high"

		result.Competition =
			"low"

		result.Recommendation =
			"enter market"

		return result
	}

	if score > 300 {

		result.Opportunity =
			"medium"

		result.Competition =
			"medium"

		result.Recommendation =
			"evaluate market"

		return result
	}

	result.Opportunity =
		"low"

	result.Competition =
		"high"

	result.Recommendation =
		"avoid market"

	return result
}
