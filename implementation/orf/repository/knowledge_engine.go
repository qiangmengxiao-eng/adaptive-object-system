package repository

import "time"

// KnowledgeEngine creates knowledge from experience.
type KnowledgeEngine struct {
	Experience *ExperienceEngine

	Store *KnowledgeStore
}

// NewKnowledgeEngine creates knowledge engine.
func NewKnowledgeEngine(
	experience *ExperienceEngine,
	store *KnowledgeStore,
) *KnowledgeEngine {

	return &KnowledgeEngine{

		Experience: experience,

		Store: store,
	}
}

// Learn creates or updates knowledge from experience.
func (k *KnowledgeEngine) Learn(
	object string,
	strategy string,
) (*Knowledge, error) {

	experience :=
		k.Experience.Analyze(
			object,
		)

	rate :=
		experience.SuccessRate()

	knowledge :=
		Knowledge{

			Object: object,

			Strategy: strategy,

			SuccessRate: rate,

			Confidence: rate,

			Preferred: rate >= 0.8,

			Global: false,

			Version: 1,

			CreatedAt: time.Now(),
		}

	if k.Store != nil {

		existing, err :=
			k.Store.Load()

		if err != nil {

			return nil, err
		}

		for _, item := range existing {

			if item.Object ==
				object &&
				item.Strategy ==
					strategy {

				knowledge.Version =
					item.Version + 1

				knowledge.Global =
					item.Global

				err =
					k.Store.Update(
						knowledge,
					)

				if err != nil {

					return nil, err
				}

				return &knowledge, nil
			}
		}

		err =
			k.Store.Append(
				knowledge,
			)

		if err != nil {

			return nil, err
		}
	}

	return &knowledge, nil
}

// Get returns object knowledge
// and global reusable knowledge.
func (k *KnowledgeEngine) Get(
	object string,
) []Knowledge {

	if k.Store == nil {

		return []Knowledge{}
	}

	list, err :=
		k.Store.Load()

	if err != nil {

		return []Knowledge{}
	}

	result :=
		make(
			[]Knowledge,
			0,
		)

	for _, item := range list {

		// Object knowledge
		// or global reusable knowledge.
		if item.Object == object ||
			item.Global {

			result =
				append(
					result,
					item,
				)
		}
	}

	return result
}

// Promote promotes successful knowledge
// into global reusable intelligence.
func (k *KnowledgeEngine) Promote(
	object string,
	strategy string,
) error {

	if k.Store == nil {

		return nil
	}

	list, err :=
		k.Store.Load()

	if err != nil {

		return err
	}

	for _, item := range list {

		if item.Object ==
			object &&
			item.Strategy ==
				strategy &&
			item.SuccessRate >= 0.9 {

			globalKnowledge :=
				item

			globalKnowledge.Object =
				""

			globalKnowledge.Global =
				true

			globalKnowledge.Version =
				item.Version + 1

			globalKnowledge.CreatedAt =
				time.Now()

			// Update existing global knowledge
			// instead of creating duplicates.
			return k.Store.Update(
				globalKnowledge,
			)
		}
	}

	return nil
}
