package repository

// MetricsEngine collects system metrics.
type MetricsEngine struct {
	Registry *Registry

	ExecutionStore *ExecutionStore

	Knowledge *KnowledgeStore

	ObservationStore *ObservationStore
}

// NewMetricsEngine creates metrics engine.
func NewMetricsEngine(
	r *Registry,
	e *ExecutionStore,
	k *KnowledgeStore,
	o *ObservationStore,
) *MetricsEngine {

	return &MetricsEngine{

		Registry: r,

		ExecutionStore: e,

		Knowledge: k,

		ObservationStore: o,
	}
}

// Collect collects system metrics.
func (m *MetricsEngine) Collect() Metrics {

	result :=
		Metrics{}

	// objects

	if m.Registry != nil {

		list, err :=
			m.Registry.List()

		if err == nil {

			result.Objects =
				len(list)
		}
	}

	// executions

	if m.ExecutionStore != nil {

		list, err :=
			m.ExecutionStore.Load()

		if err == nil {

			result.Executions =
				len(list)
		}
	}

	// knowledge

	if m.Knowledge != nil {

		list, err :=
			m.Knowledge.Load()

		if err == nil {

			result.Knowledge =
				len(list)
		}
	}

	// global success rate

	if m.ObservationStore != nil {

		list, err :=
			m.ObservationStore.Load()

		if err == nil &&
			len(list) > 0 {

			success :=
				0

			for _, item := range list {

				if item.Result ==
					"success" {

					success++
				}
			}

			result.SuccessRate =
				float64(success) /
					float64(len(list))
		}
	}

	return result
}
