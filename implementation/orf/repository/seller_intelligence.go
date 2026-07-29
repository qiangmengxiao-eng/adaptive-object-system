package repository

type SellerAnalysis struct {
	Product string `json:"product"`

	Profit ProfitAnalysis `json:"profit"`

	Market MarketStrategy `json:"market"`

	Listing *AmazonAgentResult `json:"listing"`

	Recommendation string `json:"recommendation"`
}

type SellerIntelligence struct {
	Profit *ProfitEngine

	Market *MarketStrategyEngine

	Amazon *AmazonAgent
}

func NewSellerIntelligence(
	p *ProfitEngine,
	m *MarketStrategyEngine,
	a *AmazonAgent,
) *SellerIntelligence {

	return &SellerIntelligence{

		Profit: p,

		Market: m,

		Amazon: a,
	}
}

func (s *SellerIntelligence) Analyze(
	object string,
	product string,
	cost float64,
	price float64,
	features []string,
	keywords []string,
) (*SellerAnalysis, error) {

	profit :=
		s.Profit.Analyze(
			cost,
			price,
		)

	market :=
		s.Market.Analyze(
			price,
		)

	listing, err :=
		s.Amazon.Run(
			object,
			product,
			features,
			keywords,
		)

	if err != nil {

		return nil, err
	}

	recommendation :=
		"not recommended"

	if profit.Assessment != "low" {

		recommendation =
			"recommended"
	}

	return &SellerAnalysis{

		Product: product,

		Profit: profit,

		Market: market,

		Listing: listing,

		Recommendation: recommendation,
	}, nil
}
