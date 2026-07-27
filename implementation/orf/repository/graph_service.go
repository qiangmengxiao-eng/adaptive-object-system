package repository

// GraphService manages persistent relations.
type GraphService struct {
	store *GraphStore
}

// NewGraphService creates graph service.
func NewGraphService(
	store *GraphStore,
) *GraphService {

	return &GraphService{
		store: store,
	}
}

// AddRelation persists relation.
func (s *GraphService) AddRelation(
	relation ObjectRelation,
) error {

	graph, err := s.store.Load()

	if err != nil {
		return err
	}

	if err := graph.AddRelation(
		relation,
	); err != nil {

		return err
	}

	return s.store.Save(
		graph,
	)
}

// QueryRelations finds relations.
func (s *GraphService) QueryRelations(
	from string,
) ([]ObjectRelation, error) {

	graph, err := s.store.Load()

	if err != nil {
		return nil, err
	}

	return graph.FindRelations(
		from,
	), nil
}

// List returns graph.
func (s *GraphService) List() (
	[]ObjectRelation, error,
) {

	graph, err := s.store.Load()

	if err != nil {
		return nil, err
	}

	return graph.Relations(), nil
}
