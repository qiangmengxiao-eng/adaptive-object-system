package repository

type ProductStrategyEngine struct{}

func NewProductStrategyEngine() *ProductStrategyEngine {

	return &ProductStrategyEngine{}
}

func (p *ProductStrategyEngine) Generate(
	product string,
	marketScore float64,
	competition string,
	margin float64,
) ProductStrategy {

	strategy :=
		ProductStrategy{

			Product: product,

			Decision: "review",

			Positioning: "value",

			TargetPrice: 0,

			CoreFeatures: []string{},

			KeywordStrategy: "keyword_front_loading",

			MarketingStrategy: "discovery campaign",

			Risk: "unknown",
		}

	if marketScore > 1000 &&
		margin > 40 {

		strategy.Decision =
			"launch"

		strategy.Positioning =
			"premium value"

		strategy.TargetPrice =
			49.99

		strategy.CoreFeatures =
			[]string{

				"better features",

				"quality improvement",

				"customer value",
			}

		strategy.Risk =
			"manageable"
	}

	if competition == "high" {

		strategy.Risk =
			"high competition"

		strategy.Positioning =
			"differentiated product"

		strategy.CoreFeatures =
			append(
				strategy.CoreFeatures,

				"unique selling proposition",
			)
	}

	return strategy
}
