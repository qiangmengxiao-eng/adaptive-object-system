package repository

// MetricsEngine collects system metrics.
type MetricsEngine struct {
	Registry *Registry

	Knowledge *KnowledgeStore
}

// NewMetricsEngine creates metrics engine.
func NewMetricsEngine(
	r *Registry,
	k *KnowledgeStore,
) *MetricsEngine {

	return &MetricsEngine{

		Registry: r,

		Knowledge: k,
	}
}

// Collect collects system metrics.
func (m *MetricsEngine) Collect() Metrics {

	result :=
		Metrics{}

	if m.Knowledge != nil {

		list, err :=
			m.Knowledge.Load()

		if err == nil {

			result.Knowledge =
				len(list)
		}
	}

	return result
}
