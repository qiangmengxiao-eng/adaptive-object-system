package repository

import (
	"fmt"
	"time"
)

// Observation records runtime experience.
type Observation struct {
	ID string `yaml:"id"`

	Object string `yaml:"object"`

	Event string `yaml:"event"`

	Action string `yaml:"action"`

	Result string `yaml:"result"`

	Timestamp time.Time `yaml:"timestamp"`
}

// NewObservation creates observation.
func NewObservation(
	object string,
	event string,
	action string,
	result string,
) Observation {

	return Observation{

		ID: fmt.Sprintf(
			"obs-%d",
			time.Now().UnixNano(),
		),

		Object: object,

		Event: event,

		Action: action,

		Result: result,

		Timestamp: time.Now(),
	}
}
