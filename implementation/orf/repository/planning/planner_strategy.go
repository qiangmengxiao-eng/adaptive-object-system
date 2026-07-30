package planning

type PlannerStrategy struct{}

func NewPlannerStrategy() *PlannerStrategy {

	return &PlannerStrategy{}
}

func (p *PlannerStrategy) Generate(
	goal string,
) []string {

	return []string{

		"analyze current state",

		"identify possible actions",

		"evaluate available resources",

		"select optimal strategy",

		"execute planned action",

		"measure result",

		"learn from outcome",
	}
}
