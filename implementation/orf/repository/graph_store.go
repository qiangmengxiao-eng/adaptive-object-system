package repository

import (
	"fmt"
	"path"

	"gopkg.in/yaml.v3"
)

const GraphFileName = "graph.yaml"

// GraphStore persists object relations.
type GraphStore struct {
	fs MutableGraphFS
}

// MutableGraphFS defines graph storage.
type MutableGraphFS interface {
	ReadFile(path string) ([]byte, error)
	WriteFile(path string, data []byte) error
}

// NewGraphStore creates graph storage.
func NewGraphStore(fs MutableGraphFS) *GraphStore {
	return &GraphStore{
		fs: fs,
	}
}

// graphDocument represents persisted graph.
type graphDocument struct {
	Relations []ObjectRelation `yaml:"relations"`
}

// Load loads graph.
func (s *GraphStore) Load() (*ObjectGraph, error) {

	data, err := s.fs.ReadFile(
		GraphFilePath(),
	)

	if err != nil {

		return NewObjectGraph(), nil
	}

	var doc graphDocument

	if err := yaml.Unmarshal(
		data,
		&doc,
	); err != nil {

		return nil, err
	}

	graph := NewObjectGraph()

	for _, relation := range doc.Relations {

		if err := graph.AddRelation(
			relation,
		); err != nil {

			return nil, err
		}
	}

	return graph, nil
}

// Save saves graph.
func (s *GraphStore) Save(
	graph *ObjectGraph,
) error {

	doc := graphDocument{
		Relations: graph.Relations(),
	}

	data, err := yaml.Marshal(
		doc,
	)

	if err != nil {
		return err
	}

	return s.fs.WriteFile(
		GraphFilePath(),
		data,
	)
}

// GraphFilePath returns graph path.
func GraphFilePath() string {

	return path.Join(
		"/",
		GraphFileName,
	)
}

func validateGraphStore(
	s *GraphStore,
) error {

	if s == nil || s.fs == nil {
		return fmt.Errorf(
			"graph store is nil",
		)
	}

	return nil
}
