package repository

// Registry manages objects.
type Registry struct {
	repository *Repository
}

// NewRegistry creates a registry.
func NewRegistry(repository *Repository) *Registry {
	return &Registry{
		repository: repository,
	}
}

// Register creates an object and metadata.
func (r *Registry) Register(
	name string,
	definition []byte,
	metadata *ObjectMetadata,
) error {
	if err := r.repository.CreateObject(name, definition); err != nil {
		return err
	}

	return r.repository.WriteObjectMetadata(name, metadata)
}

// Get returns an object.
func (r *Registry) Get(name string) (*Object, error) {
	return r.repository.LoadObject(name)
}

// List returns object names.
func (r *Registry) List() ([]string, error) {
	return r.repository.ListObjects()
}
