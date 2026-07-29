package repository

import "time"

// Reflection represents system self evaluation.
type Reflection struct {
	Object string `json:"object" yaml:"object"`

	Events int `json:"events" yaml:"events"`

	Success int `json:"success" yaml:"success"`

	Failure int `json:"failure" yaml:"failure"`

	SuccessRate float64 `json:"success_rate" yaml:"success_rate"`

	Assessment string `json:"assessment" yaml:"assessment"`

	CreatedAt time.Time `json:"created_at" yaml:"created_at"`
}
