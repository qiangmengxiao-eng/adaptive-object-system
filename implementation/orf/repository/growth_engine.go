package repository

type GrowthResult struct {
	Competitor CompetitorAnalysis `json:"competitor"`

	SEO SEOAnalysis `json:"seo"`

	PPC PPCStrategy `json:"ppc"`

	Recommendation string `json:"recommendation"`
}

type GrowthEngine struct {
	Competitor *CompetitorEngine

	SEO *SEOEngine

	PPC *PPCEngine
}

func NewGrowthEngine(
	c *CompetitorEngine,
	s *SEOEngine,
	p *PPCEngine,
) *GrowthEngine {

	return &GrowthEngine{

		Competitor: c,

		SEO: s,

		PPC: p,
	}
}

func (g *GrowthEngine) Analyze(
	product string,
	competitors []Competitor,
	keywords []string,
) GrowthResult {

	return GrowthResult{

		Competitor: g.Competitor.Analyze(
			product,
			competitors,
		),

		SEO: g.SEO.Analyze(
			keywords,
		),

		PPC: g.PPC.Generate(
			keywords,
		),

		Recommendation: "growth strategy generated",
	}
}
