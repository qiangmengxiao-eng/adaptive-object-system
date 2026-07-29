package repository

type ProductAnalysis struct {
	Product string `json:"product"`

	Category string `json:"category"`

	TargetCustomer string `json:"target_customer"`

	Position string `json:"position"`

	KeySellingPoints []string `json:"key_selling_points"`
}

type ProductAnalysisEngine struct{}

func NewProductAnalysisEngine() *ProductAnalysisEngine {

	return &ProductAnalysisEngine{}
}

func (p *ProductAnalysisEngine) Analyze(
	product string,
	features []string,
) ProductAnalysis {

	category :=
		"general"

	if len(features) > 0 {

		category =
			"electronics"
	}

	return ProductAnalysis{

		Product: product,

		Category: category,

		TargetCustomer: "online shoppers",

		Position: "quality value",

		KeySellingPoints: features,
	}
}
