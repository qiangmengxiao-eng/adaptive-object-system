package repository

// ObjectQuery defines object search conditions.
type ObjectQuery struct {
	Name string
	Type string
}

// QueryObjects searches objects by conditions.
func (r *Repository) QueryObjects(query ObjectQuery) ([]string, error) {
	objects, err := r.ListObjects()
	if err != nil {
		return nil, err
	}

	result := make([]string, 0)

	for _, name := range objects {
		if query.Name != "" && name != query.Name {
			continue
		}

		if query.Type != "" {
			object, err := r.ReadObjectDefinition(name)
			if err != nil {
				continue
			}

			if object.Type != query.Type {
				continue
			}
		}

		result = append(result, name)
	}

	return result, nil
}
