package repository

import (
	"gopkg.in/yaml.v3"
)

const ObservationFileName = "observations.yaml"

type ObservationStore struct {
	fs MutableGraphFS
}

type observationDocument struct {
	Observations []Observation `yaml:"observations"`
}

func NewObservationStore(
	fs MutableGraphFS,
) *ObservationStore {

	return &ObservationStore{

		fs: fs,
	}
}

func (s *ObservationStore) Load() ([]Observation, error) {

	data, err :=
		s.fs.ReadFile(
			ObservationFilePath(),
		)

	if err != nil {

		return []Observation{}, nil
	}

	var doc observationDocument

	err =
		yaml.Unmarshal(
			data,
			&doc,
		)

	if err != nil {

		return nil, err
	}

	return doc.Observations, nil
}

func (s *ObservationStore) Append(
	observation Observation,
) error {

	observations, err :=
		s.Load()

	if err != nil {

		return err
	}

	observations =
		append(
			observations,
			observation,
		)

	doc :=
		observationDocument{

			Observations: observations,
		}

	data, err :=
		yaml.Marshal(
			doc,
		)

	if err != nil {

		return err
	}

	return s.fs.WriteFile(
		ObservationFilePath(),
		data,
	)
}

func ObservationFilePath() string {

	return "/observations.yaml"
}
