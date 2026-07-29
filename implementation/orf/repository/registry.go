package repository

// Registry manages objects.
type Registry struct {
	repository *Repository
}

// NewRegistry creates registry.
func NewRegistry(
	repository *Repository,
) *Registry {

	return &Registry{

		repository: repository,
	}
}

// Register creates object and metadata.
func (r *Registry) Register(
	name string,
	definition []byte,
	metadata *ObjectMetadata,
) error {

	if err :=
		r.repository.CreateObject(
			name,
			definition,
		); err != nil {

		return err
	}

	if err :=
		r.repository.WriteObjectMetadata(
			name,
			metadata,
		); err != nil {

		return err
	}

	return nil
}

// Get returns object.
func (r *Registry) Get(
	name string,
) (*Object, error) {

	return r.repository.LoadObject(
		name,
	)
}

// List returns objects.
func (r *Registry) List() (
	[]string, error,
) {

	return r.repository.ListObjects()
}
