package repository

import "gopkg.in/yaml.v3"

type CollaborationStore struct {
	fs MutableGraphFS
}

type collaborationDocument struct {
	Collaborations []Collaboration `yaml:"collaborations"`
}

func NewCollaborationStore(
	fs MutableGraphFS,
) *CollaborationStore {

	return &CollaborationStore{
		fs: fs,
	}
}

func (s *CollaborationStore) Load() (
	[]Collaboration, error,
) {

	data, err :=
		s.fs.ReadFile(
			"/collaboration.yaml",
		)

	if err != nil {

		return []Collaboration{}, nil
	}

	var doc collaborationDocument

	err =
		yaml.Unmarshal(
			data,
			&doc,
		)

	return doc.Collaborations, err
}

func (s *CollaborationStore) Append(
	item Collaboration,
) error {

	list, _ :=
		s.Load()

	list =
		append(
			list,
			item,
		)

	data, err :=
		yaml.Marshal(
			collaborationDocument{
				Collaborations: list,
			},
		)

	if err != nil {

		return err
	}

	return s.fs.WriteFile(
		"/collaboration.yaml",
		data,
	)
}
