package repository

type ExperienceEngine struct {
	observations *ObservationStore
}

func NewExperienceEngine(
	store *ObservationStore,
) *ExperienceEngine {

	return &ExperienceEngine{

		observations: store,
	}
}

func (e *ExperienceEngine) Analyze(
	object string,
) Experience {

	list, err :=
		e.observations.Load()

	if err != nil {

		return Experience{
			Object: object,
		}
	}

	result :=
		Experience{

			Object: object,
		}

	for _, item := range list {

		if item.Object != object {

			continue
		}

		result.Events++

		if item.Result == "success" {

			result.Success++

		} else {

			result.Failure++
		}
	}

	return result
}
