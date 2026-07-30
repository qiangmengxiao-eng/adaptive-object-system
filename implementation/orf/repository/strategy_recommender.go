package repository

type StrategyRecommendation struct {
	Strategy string `json:"strategy"`

	Confidence float64 `json:"confidence"`

	Reason string `json:"reason"`
}

type StrategyRecommender struct {
	Store *StrategyKnowledgeStore
}

func NewStrategyRecommender(
	store *StrategyKnowledgeStore,
) *StrategyRecommender {

	return &StrategyRecommender{
		Store: store,
	}
}

func (r *StrategyRecommender) Recommend(
	category string,
) StrategyRecommendation {

	list, err :=
		r.Store.Load()

	if err != nil {

		return StrategyRecommendation{}
	}

	for _, item := range list {

		if item.Category ==
			category &&
			item.Preferred {

			return StrategyRecommendation{

				Strategy: item.Strategy,

				Confidence: item.Confidence,

				Reason: "learned from successful products",
			}
		}
	}

	return StrategyRecommendation{

		Strategy: "no previous strategy",

		Confidence: 0,

		Reason: "insufficient knowledge",
	}
}
