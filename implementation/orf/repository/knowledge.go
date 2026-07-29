package repository

import "time"

// Knowledge represents learned strategy knowledge.
//
// Knowledge can belong to a single object
// or become global reusable intelligence.
type Knowledge struct {

	// Object identifies the source object.
	//
	// Empty object means global knowledge.
	Object string `json:"object" yaml:"object"`

	// Strategy represents learned execution strategy.
	Strategy string `json:"strategy" yaml:"strategy"`

	// SuccessRate represents historical success ratio.
	SuccessRate float64 `json:"success_rate" yaml:"success_rate"`

	// Confidence represents confidence of this knowledge.
	Confidence float64 `json:"confidence" yaml:"confidence"`

	// Preferred indicates this strategy
	// should be considered first.
	Preferred bool `json:"preferred" yaml:"preferred"`

	// Global indicates this knowledge
	// can be reused by other objects.
	Global bool `json:"global" yaml:"global"`

	// Version represents knowledge evolution version.
	Version int `json:"version" yaml:"version"`

	// CreatedAt records knowledge creation time.
	CreatedAt time.Time `json:"created_at" yaml:"created_at"`
}
