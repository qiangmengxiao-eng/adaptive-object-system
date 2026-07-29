package repository

import "time"

type ReflectionEngine struct {
	Experience *ExperienceEngine

	Store *ReflectionStore
}

func NewReflectionEngine(
	experience *ExperienceEngine,
	store *ReflectionStore,
) *ReflectionEngine {

	return &ReflectionEngine{

		Experience: experience,

		Store: store,
	}
}

func (r *ReflectionEngine) Reflect(
	object string,
) Reflection {

	experience :=
		r.Experience.Analyze(
			object,
		)

	rate := float64(0)

	if experience.Events > 0 {

		rate =
			float64(experience.Success) /
				float64(experience.Events)
	}

	assessment :=
		"needs improvement"

	if rate >= 0.8 {

		assessment =
			"performance acceptable"
	}

	reflection :=
		Reflection{

			Object: object,

			Events: experience.Events,

			Success: experience.Success,

			Failure: experience.Failure,

			SuccessRate: rate,

			Assessment: assessment,

			CreatedAt: time.Now(),
		}

	if r.Store != nil {

		_ =
			r.Store.Append(
				reflection,
			)
	}

	return reflection
}

func (r *ReflectionEngine) Latest(
	object string,
) *Reflection {

	list, err :=
		r.Store.Load()

	if err != nil {

		return nil
	}

	for i := len(list) - 1; i >= 0; i-- {

		if list[i].Object == object {

			return &list[i]
		}
	}

	return nil
}
