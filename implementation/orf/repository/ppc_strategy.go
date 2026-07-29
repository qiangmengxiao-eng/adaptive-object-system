package repository

type PPCStrategy struct {
	Campaign string `json:"campaign"`

	Keywords []string `json:"keywords"`

	Bidding string `json:"bidding"`

	Recommendation string `json:"recommendation"`
}

type PPCEngine struct{}

func NewPPCEngine() *PPCEngine {

	return &PPCEngine{}
}

func (p *PPCEngine) Generate(
	keywords []string,
) PPCStrategy {

	return PPCStrategy{

		Campaign: "auto-keyword-campaign",

		Keywords: keywords,

		Bidding: "dynamic",

		Recommendation: "start with discovery campaign",
	}
}
