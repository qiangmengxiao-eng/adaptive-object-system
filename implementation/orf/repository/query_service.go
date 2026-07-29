package repository

// QueryService provides object queries.
type QueryService struct {
	registry *Registry
}

// NewQueryService creates query service.
func NewQueryService(
	registry *Registry,
) *QueryService {

	return &QueryService{

		registry: registry,
	}
}

// ListObjects returns all objects.
func (q *QueryService) ListObjects() ([]string, error) {

	return q.registry.List()
}
