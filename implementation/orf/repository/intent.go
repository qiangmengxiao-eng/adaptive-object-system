package repository

import "time"

// Intent represents object intention.
type Intent struct {
	Object string `json:"object"`

	Goal string `json:"goal"`

	Name string `json:"name"`

	Purpose string `json:"purpose"`

	Status string `json:"status"`

	Version int `json:"version"`

	CreatedAt time.Time `json:"created_at"`
}

// IntentEngine manages object intents.
type IntentEngine struct {
	intents []Intent
}

// NewIntentEngine creates intent engine.
func NewIntentEngine() *IntentEngine {

	return &IntentEngine{

		intents: make(
			[]Intent,
			0,
		),
	}
}

// Create creates an intent.
func (e *IntentEngine) Create(
	object string,
	goal string,
	name string,
	purpose string,
) Intent {

	intent :=
		Intent{

			Object: object,

			Goal: goal,

			Name: name,

			Purpose: purpose,

			Status: "active",

			Version: 1,

			CreatedAt: time.Now(),
		}

	e.intents =
		append(
			e.intents,
			intent,
		)

	return intent
}

// Get returns object intents.
func (e *IntentEngine) Get(
	object string,
) []Intent {

	result :=
		make(
			[]Intent,
			0,
		)

	for _, intent := range e.intents {

		if intent.Object ==
			object {

			result =
				append(
					result,
					intent,
				)
		}
	}

	return result
}
