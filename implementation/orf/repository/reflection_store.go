package repository

import (
	"gopkg.in/yaml.v3"
)

const ReflectionFileName = "reflections.yaml"

type ReflectionStore struct {
	fs MutableGraphFS
}

type reflectionDocument struct {
	Reflections []Reflection `yaml:"reflections"`
}

func NewReflectionStore(
	fs MutableGraphFS,
) *ReflectionStore {

	return &ReflectionStore{
		fs: fs,
	}
}

func (s *ReflectionStore) Load() (
	[]Reflection,
	error,
) {

	data, err :=
		s.fs.ReadFile(
			ReflectionFilePath(),
		)

	if err != nil {

		return []Reflection{}, nil
	}

	var doc reflectionDocument

	err =
		yaml.Unmarshal(
			data,
			&doc,
		)

	if err != nil {

		return nil, err
	}

	return doc.Reflections, nil
}

func (s *ReflectionStore) Append(
	reflection Reflection,
) error {

	list, err :=
		s.Load()

	if err != nil {

		return err
	}

	list =
		append(
			list,
			reflection,
		)

	data, err :=
		yaml.Marshal(
			reflectionDocument{
				Reflections: list,
			},
		)

	if err != nil {

		return err
	}

	return s.fs.WriteFile(
		ReflectionFilePath(),
		data,
	)
}

func ReflectionFilePath() string {

	return "/reflections.yaml"
}
