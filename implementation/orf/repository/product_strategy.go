package repository

type ProductStrategy struct {
	Product string `json:"product"`

	Decision string `json:"decision"`

	Positioning string `json:"positioning"`

	TargetPrice float64 `json:"target_price"`

	CoreFeatures []string `json:"core_features"`

	KeywordStrategy string `json:"keyword_strategy"`

	MarketingStrategy string `json:"marketing_strategy"`

	Risk string `json:"risk"`
}
