package repository

type AmazonAgentResult struct {
	Analysis ProductAnalysis `json:"analysis"`

	Keywords KeywordStrategy `json:"keywords"`

	Listing AmazonListing `json:"listing"`

	Quality ListingQuality `json:"quality"`

	Recommendation string `json:"recommendation"`
}

type AmazonAgent struct {
	ProductAnalysis *ProductAnalysisEngine

	KeywordStrategy *KeywordStrategyEngine

	Listing *ListingEngine

	Quality *ListingQualityEngine
}

func NewAmazonAgent(
	p *ProductAnalysisEngine,
	k *KeywordStrategyEngine,
	l *ListingEngine,
	q *ListingQualityEngine,
) *AmazonAgent {

	return &AmazonAgent{

		ProductAnalysis: p,

		KeywordStrategy: k,

		Listing: l,

		Quality: q,
	}
}

func (a *AmazonAgent) Run(
	object string,
	product string,
	features []string,
	keywords []string,
) (*AmazonAgentResult, error) {

	analysis :=
		a.ProductAnalysis.Analyze(
			product,
			features,
		)

	strategy :=
		a.KeywordStrategy.Generate(
			keywords,
		)

	listing, err :=
		a.Listing.Generate(
			object,
			product,
			features,
			keywords,
		)

	if err != nil {

		return nil, err
	}

	quality :=
		a.Quality.Analyze(
			*listing,
		)

	recommendation :=
		"ready"

	if quality.OverallScore < 80 {

		recommendation =
			"needs optimization"
	}

	return &AmazonAgentResult{

		Analysis: analysis,

		Keywords: strategy,

		Listing: *listing,

		Quality: quality,

		Recommendation: recommendation,
	}, nil
}
