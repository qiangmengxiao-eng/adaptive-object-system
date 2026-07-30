package repository

import "time"

type StrategyKnowledge struct {
	ID string `json:"id" yaml:"id"`

	Object string `json:"object" yaml:"object"`

	Strategy string `json:"strategy" yaml:"strategy"`

	Category string `json:"category" yaml:"category"`

	SuccessRate float64 `json:"success_rate" yaml:"success_rate"`

	Confidence float64 `json:"confidence" yaml:"confidence"`

	Preferred bool `json:"preferred" yaml:"preferred"`

	Version int `json:"version" yaml:"version"`

	CreatedAt time.Time `json:"created_at" yaml:"created_at"`
}
