package repository

import "time"

// RuntimeObject represents a running object instance.
type RuntimeObject struct {
	ID string `yaml:"id"`

	Name string `yaml:"name"`

	Definition ObjectDefinition `yaml:"definition"`

	State RuntimeState `yaml:"state"`

	Events []ObjectEvent `yaml:"events"`

	Behaviors []string `yaml:"behaviors"`

	Relations []ObjectRelation `yaml:"relations"`

	CreatedAt time.Time `yaml:"created_at"`

	UpdatedAt time.Time `yaml:"updated_at"`
}

// NewRuntimeObject creates runtime object.
func NewRuntimeObject(
	definition ObjectDefinition,
) *RuntimeObject {

	now :=
		time.Now()

	return &RuntimeObject{

		ID: definition.ID,

		Name: definition.Name,

		Definition: definition,

		State: NewRuntimeState(
			"lifecycle",
			"created",
		),

		Events: make(
			[]ObjectEvent,
			0,
		),

		Behaviors: make(
			[]string,
			0,
		),

		Relations: make(
			[]ObjectRelation,
			0,
		),

		CreatedAt: now,

		UpdatedAt: now,
	}
}

// AddEvent appends event.
func (r *RuntimeObject) AddEvent(
	event ObjectEvent,
) {

	if r == nil {

		return
	}

	r.Events =
		append(
			r.Events,
			event,
		)

	r.UpdatedAt =
		time.Now()
}
