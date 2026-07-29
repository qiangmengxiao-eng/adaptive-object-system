package repository

import (
	"fmt"
	"time"
)

type ObjectEvent struct {
	ID string `yaml:"id"`

	Type string `yaml:"type"`

	Object string `yaml:"object"`

	Action string `yaml:"action"`

	Data string `yaml:"data,omitempty"`

	Version int `yaml:"version"`

	Timestamp time.Time `yaml:"timestamp"`
}

func NewObjectEvent(
	eventType string,
	object string,
	action string,
	data string,
) ObjectEvent {

	return ObjectEvent{

		ID: fmt.Sprintf(
			"evt-%d",
			time.Now().UnixNano(),
		),

		Type: eventType,

		Object: object,

		Action: action,

		Data: data,

		Version: 1,

		Timestamp: time.Now(),
	}
}

func (e *ObjectEvent) Validate() error {

	if e == nil {
		return fmt.Errorf(
			"event is nil",
		)
	}

	if e.Type == "" {
		return fmt.Errorf(
			"event type required",
		)
	}

	if e.Object == "" {
		return fmt.Errorf(
			"event object required",
		)
	}

	if e.Version == 0 {

		e.Version = 1
	}

	if e.Timestamp.IsZero() {

		e.Timestamp = time.Now()
	}

	return nil
}
