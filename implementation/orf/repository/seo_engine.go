package repository

type SEOAnalysis struct {
	PrimaryKeywords []string `json:"primary_keywords"`

	SecondaryKeywords []string `json:"secondary_keywords"`

	Recommendation string `json:"recommendation"`
}

type SEOEngine struct{}

func NewSEOEngine() *SEOEngine {

	return &SEOEngine{}
}

func (s *SEOEngine) Analyze(
	keywords []string,
) SEOAnalysis {

	primary :=
		[]string{}

	secondary :=
		[]string{}

	for i, k := range keywords {

		if i < 2 {

			primary =
				append(
					primary,
					k,
				)

		} else {

			secondary =
				append(
					secondary,
					k,
				)
		}
	}

	return SEOAnalysis{

		PrimaryKeywords: primary,

		SecondaryKeywords: secondary,

		Recommendation: "keyword optimization recommended",
	}
}
