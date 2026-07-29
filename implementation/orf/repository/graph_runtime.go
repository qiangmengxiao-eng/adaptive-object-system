package repository

// RuntimeGraph represents graph information
// attached to runtime object.
type RuntimeGraph struct {
	Object string `yaml:"object"`

	State string `yaml:"state"`

	Relations []ObjectRelation `yaml:"relations"`

	Events []ObjectEvent `yaml:"events"`
}

// GraphRuntimeService connects graph and runtime.
type GraphRuntimeService struct {
	graph *GraphService

	runtime *RuntimeEngine
}

// NewGraphRuntimeService creates service.
func NewGraphRuntimeService(
	graph *GraphService,
	runtime *RuntimeEngine,
) *GraphRuntimeService {

	return &GraphRuntimeService{

		graph: graph,

		runtime: runtime,
	}
}

// Inspect returns runtime graph.
func (s *GraphRuntimeService) Inspect(
	name string,
) RuntimeGraph {

	result :=
		RuntimeGraph{

			Object: name,

			Relations: make(
				[]ObjectRelation,
				0,
			),

			Events: make(
				[]ObjectEvent,
				0,
			),
		}

	// load runtime

	if s.runtime != nil {

		if object, ok :=
			s.runtime.Get(
				name,
			); ok {

			result.State =
				object.State.Status

			result.Events =
				object.Events
		}
	}

	// load graph relations

	if s.graph != nil {

		relations, err :=
			s.graph.QueryRelations(
				name,
			)

		if err == nil {

			result.Relations =
				relations
		}
	}

	return result
}
