package repository

type KnowledgeLearningEngine struct {
	Store *StrategyKnowledgeStore
}

func NewKnowledgeLearningEngine(
	store *StrategyKnowledgeStore,
) *KnowledgeLearningEngine {

	return &KnowledgeLearningEngine{

		Store: store,
	}
}

func (k *KnowledgeLearningEngine) Learn(
	object string,
	category string,
	strategy string,
	successRate float64,
	confidence float64,
) error {

	item :=
		StrategyKnowledge{

			Object: object,

			Strategy: strategy,

			Category: category,

			SuccessRate: successRate,

			Confidence: confidence,

			Preferred: confidence > 0.8,

			Version: 1,
		}

	return k.Store.Append(
		item,
	)
}
