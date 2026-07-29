package repository

import (
	"gopkg.in/yaml.v3"
)

const RuntimeFileName = "runtime.yaml"

// RuntimeStore persists runtime objects.
type RuntimeStore struct {
	fs Storage
}

// NewRuntimeStore creates runtime store.
func NewRuntimeStore(
	fs Storage,
) *RuntimeStore {

	return &RuntimeStore{
		fs: fs,
	}
}

// LoadAll loads all runtime objects.
func (s *RuntimeStore) LoadAll() (
	[]*RuntimeObject,
	error,
) {

	data, err :=
		s.fs.ReadFile(
			"/" + RuntimeFileName,
		)

	if err != nil {

		return []*RuntimeObject{}, nil
	}

	var document struct {
		Objects []*RuntimeObject `yaml:"objects"`
	}

	err =
		yaml.Unmarshal(
			data,
			&document,
		)

	if err != nil {

		return nil, err
	}

	return document.Objects, nil
}

// SaveAll persists runtime objects.
func (s *RuntimeStore) SaveAll(
	objects []*RuntimeObject,
) error {

	document :=
		struct {
			Objects []*RuntimeObject `yaml:"objects"`
		}{
			Objects: objects,
		}

	data, err :=
		yaml.Marshal(
			document,
		)

	if err != nil {

		return err
	}

	return s.fs.WriteFile(
		"/"+RuntimeFileName,
		data,
	)
}
