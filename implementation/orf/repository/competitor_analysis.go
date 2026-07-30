package repository

type Competitor struct {
	Name string `json:"name"`

	Price float64 `json:"price"`

	Rating float64 `json:"rating"`

	Reviews int `json:"reviews"`

	Keywords []string `json:"keywords"`
}

type CompetitorAnalysis struct {
	Product string `json:"product"`

	CompetitionLevel string `json:"competition_level"`

	Opportunity string `json:"opportunity"`

	Gap string `json:"gap"`

	AveragePrice float64 `json:"average_price"`

	AverageRating float64 `json:"average_rating"`

	AverageReviews float64 `json:"average_reviews"`

	Differentiation string `json:"differentiation"`

	Strategy string `json:"strategy"`
}

type CompetitorEngine struct{}

func NewCompetitorEngine() *CompetitorEngine {

	return &CompetitorEngine{}
}

func (c *CompetitorEngine) Analyze(
	product string,
	competitors []Competitor,
) CompetitorAnalysis {

	averagePrice := 0.0

	averageRating := 0.0

	averageReviews := 0.0

	for _, item := range competitors {

		averagePrice += item.Price

		averageRating += item.Rating

		averageReviews += float64(item.Reviews)
	}

	if len(competitors) > 0 {

		count :=
			float64(len(competitors))

		averagePrice =
			averagePrice / count

		averageRating =
			averageRating / count

		averageReviews =
			averageReviews / count
	}

	level :=
		"low"

	if len(competitors) >= 5 {

		level =
			"high"
	}

	opportunity :=
		"standard"

	gap :=
		"unknown"

	differentiation :=
		"improve features and value"

	strategy :=
		"differentiate product"

	if len(competitors) > 0 {

		gap =
			"feature differentiation"

		opportunity =
			"improve positioning"

		if averageRating >= 4.5 &&
			averageReviews > 1000 {

			level =
				"high"

			gap =
				"strong competitors"

			differentiation =
				"better features and customer value"

			strategy =
				"enter with differentiated product"
		}
	}

	return CompetitorAnalysis{

		Product: product,

		CompetitionLevel: level,

		Opportunity: opportunity,

		Gap: gap,

		AveragePrice: averagePrice,

		AverageRating: averageRating,

		AverageReviews: averageReviews,

		Differentiation: differentiation,

		Strategy: strategy,
	}
}
