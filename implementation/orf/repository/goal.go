package repository

import "time"

// Goal represents object objective.
type Goal struct {
	Object string `json:"object"`

	Name string `json:"name"`

	Description string `json:"description"`

	Status string `json:"status"`

	Version int `json:"version"`

	CreatedAt time.Time `json:"created_at"`
}

// GoalEngine manages object goals.
type GoalEngine struct {
	goals []Goal
}

// NewGoalEngine creates goal engine.
func NewGoalEngine() *GoalEngine {

	return &GoalEngine{

		goals: make(
			[]Goal,
			0,
		),
	}
}

// Create creates a goal.
func (e *GoalEngine) Create(
	object string,
	name string,
	description string,
) Goal {

	goal :=
		Goal{

			Object: object,

			Name: name,

			Description: description,

			Status: "active",

			Version: 1,

			CreatedAt: time.Now(),
		}

	e.goals =
		append(
			e.goals,
			goal,
		)

	return goal
}

// Get returns object goal.
func (e *GoalEngine) Get(
	object string,
) []Goal {

	result :=
		make(
			[]Goal,
			0,
		)

	for _, goal := range e.goals {

		if goal.Object == object {

			result =
				append(
					result,
					goal,
				)
		}
	}

	return result
}
