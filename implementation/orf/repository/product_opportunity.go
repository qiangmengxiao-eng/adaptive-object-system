package repository

import "time"

type ProductOpportunityEngine struct {
	Market *MarketEngine
}

func NewProductOpportunityEngine(
	market *MarketEngine,
) *ProductOpportunityEngine {

	return &ProductOpportunityEngine{

		Market: market,
	}
}

func (p *ProductOpportunityEngine) Discover(
	keyword string,
	searchVolume int,
	competition int,
) MarketOpportunity {

	result :=
		p.Market.Analyze(
			searchVolume,
			competition,
		)

	return MarketOpportunity{

		Object: "market-analysis",

		Keyword: keyword,

		SearchVolume: searchVolume,

		Competition: competition,

		Opportunity: result.Opportunity,

		Recommendation: result.Recommendation,

		Score: result.Score,

		Version: 1,

		CreatedAt: time.Now(),
	}
}
