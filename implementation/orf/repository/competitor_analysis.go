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
}

type CompetitorEngine struct{}

func NewCompetitorEngine() *CompetitorEngine {

	return &CompetitorEngine{}
}

func (c *CompetitorEngine) Analyze(
	product string,
	competitors []Competitor,
) CompetitorAnalysis {

	level :=
		"low"

	if len(competitors) > 5 {

		level =
			"high"
	}

	opportunity :=
		"standard"

	gap :=
		"unknown"

	if len(competitors) > 0 {

		gap =
			"feature differentiation"

		opportunity =
			"improve positioning"
	}

	return CompetitorAnalysis{

		Product: product,

		CompetitionLevel: level,

		Opportunity: opportunity,

		Gap: gap,
	}
}
