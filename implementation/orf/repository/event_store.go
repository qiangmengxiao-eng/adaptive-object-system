package repository

import (
	"path"

	"gopkg.in/yaml.v3"
)

const EventFileName = "events.yaml"

type EventStore struct {
	fs MutableGraphFS
}

func NewEventStore(
	fs MutableGraphFS,
) *EventStore {

	return &EventStore{
		fs: fs,
	}
}

type eventDocument struct {
	Events []ObjectEvent `yaml:"events"`
}

func (s *EventStore) Load() (
	[]ObjectEvent,
	error,
) {

	data, err := s.fs.ReadFile(
		EventFilePath(),
	)

	if err != nil {
		return []ObjectEvent{}, nil
	}

	var doc eventDocument

	if err := yaml.Unmarshal(
		data,
		&doc,
	); err != nil {
		return nil, err
	}

	return doc.Events, nil
}

func (s *EventStore) Save(
	events []ObjectEvent,
) error {

	data, err := yaml.Marshal(
		eventDocument{
			Events: events,
		},
	)

	if err != nil {
		return err
	}

	return s.fs.WriteFile(
		EventFilePath(),
		data,
	)
}

func EventFilePath() string {

	return path.Join(
		"/",
		EventFileName,
	)
}
