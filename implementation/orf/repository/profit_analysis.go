package repository

type ProfitAnalysis struct {
	Cost float64 `json:"cost"`

	Price float64 `json:"price"`

	GrossProfit float64 `json:"gross_profit"`

	Margin float64 `json:"margin"`

	Assessment string `json:"assessment"`
}

type ProfitEngine struct{}

func NewProfitEngine() *ProfitEngine {

	return &ProfitEngine{}
}

func (p *ProfitEngine) Analyze(
	cost float64,
	price float64,
) ProfitAnalysis {

	profit :=
		price - cost

	margin :=
		(profit / price) * 100

	assessment :=
		"low"

	if margin >= 30 {

		assessment =
			"healthy"
	}

	if margin >= 50 {

		assessment =
			"excellent"
	}

	return ProfitAnalysis{

		Cost: cost,

		Price: price,

		GrossProfit: profit,

		Margin: margin,

		Assessment: assessment,
	}
}
