package repository

import (
	"time"
)

// Adaptation represents a system adaptation result.
type Adaptation struct {
	Object string `json:"object"`

	Trigger string `json:"trigger"`

	Previous string `json:"previous"`

	NewBehavior string `json:"new_behavior"`

	Reason string `json:"reason"`

	Confidence float64 `json:"confidence"`

	Version int `json:"version"`

	Timestamp time.Time `json:"timestamp"`
}

// AdaptationEngine evolves object behavior.
type AdaptationEngine struct {
	Experience *ExperienceEngine

	Behavior *BehaviorEngine

	Knowledge *KnowledgeEngine

	history []Adaptation
}

// NewAdaptationEngine creates adaptation engine.
func NewAdaptationEngine(
	experience *ExperienceEngine,
	behavior *BehaviorEngine,
) *AdaptationEngine {

	return &AdaptationEngine{

		Experience: experience,

		Behavior: behavior,

		history: make(
			[]Adaptation,
			0,
		),
	}
}

// AttachKnowledge attaches knowledge.
func (e *AdaptationEngine) AttachKnowledge(
	knowledge *KnowledgeEngine,
) {

	e.Knowledge =
		knowledge
}

// Adapt applies adaptation decision.
func (e *AdaptationEngine) Adapt(
	object string,
	decision *Decision,
) (
	*Adaptation,
	error,
) {

	if decision == nil {

		return nil, nil
	}

	if decision.Type != "adapt" {

		return nil, nil
	}

	newBehavior :=
		"retry"

	// Knowledge driven behavior evolution.
	if e.Knowledge != nil {

		list :=
			e.Knowledge.Get(
				object,
			)

		for _, item := range list {

			if item.Preferred {

				newBehavior =
					item.Strategy

				break
			}
		}
	}

	adaptation :=
		Adaptation{

			Object: object,

			Trigger: decision.Type,

			Previous: "default",

			NewBehavior: newBehavior,

			Reason: decision.Reason,

			Confidence: decision.Confidence,

			Version: len(e.history) + 1,

			Timestamp: time.Now(),
		}

	if e.Behavior != nil {

		_ =
			e.Behavior.Register(
				ObjectBehavior{

					Name: newBehavior,

					Action: "execute",
				},
			)
	}

	e.history =
		append(
			e.history,
			adaptation,
		)

	return &adaptation, nil
}

// List returns adaptation history.
func (e *AdaptationEngine) List(
	object string,
) []Adaptation {

	result :=
		make(
			[]Adaptation,
			0,
		)

	for _, item := range e.history {

		if item.Object ==
			object {

			result =
				append(
					result,
					item,
				)
		}
	}

	return result
}
