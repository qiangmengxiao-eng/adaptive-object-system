package repository

type MarketStrategy struct {
	Position string `json:"position"`

	Strategy string `json:"strategy"`

	Reason string `json:"reason"`
}

type MarketStrategyEngine struct{}

func NewMarketStrategyEngine() *MarketStrategyEngine {

	return &MarketStrategyEngine{}
}

func (m *MarketStrategyEngine) Analyze(
	price float64,
) MarketStrategy {

	if price >= 50 {

		return MarketStrategy{

			Position: "premium",

			Strategy: "benefit_focused",

			Reason: "higher price positioning",
		}
	}

	if price >= 20 {

		return MarketStrategy{

			Position: "value",

			Strategy: "feature_focused",

			Reason: "balanced pricing",
		}
	}

	return MarketStrategy{

		Position: "budget",

		Strategy: "price_focused",

		Reason: "competitive pricing",
	}
}
