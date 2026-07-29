package repository

import (
	"path"

	"gopkg.in/yaml.v3"
)

const ExecutionFileName = "executions.yaml"

// ExecutionStore persists execution records.
type ExecutionStore struct {
	fs MutableGraphFS
}

// NewExecutionStore creates execution store.
func NewExecutionStore(
	fs MutableGraphFS,
) *ExecutionStore {

	return &ExecutionStore{
		fs: fs,
	}
}

type executionDocument struct {
	Executions []ExecutionRecord `yaml:"executions"`
}

// Load reads execution records.
func (s *ExecutionStore) Load() (
	[]ExecutionRecord,
	error,
) {

	data, err :=
		s.fs.ReadFile(
			ExecutionFilePath(),
		)

	if err != nil {

		return []ExecutionRecord{}, nil
	}

	var doc executionDocument

	if err :=
		yaml.Unmarshal(
			data,
			&doc,
		); err != nil {

		return nil, err
	}

	return doc.Executions, nil
}

// Save stores execution records.
func (s *ExecutionStore) Save(
	records []ExecutionRecord,
) error {

	data, err :=
		yaml.Marshal(
			executionDocument{
				Executions: records,
			},
		)

	if err != nil {

		return err
	}

	return s.fs.WriteFile(
		ExecutionFilePath(),
		data,
	)
}

// ExecutionFilePath returns execution storage path.
func ExecutionFilePath() string {

	return path.Join(
		"/",
		ExecutionFileName,
	)
}
