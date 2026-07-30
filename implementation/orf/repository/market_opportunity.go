package repository

import "time"

type MarketOpportunity struct {
	Object string `json:"object" yaml:"object"`

	Keyword string `json:"keyword" yaml:"keyword"`

	SearchVolume int `json:"search_volume" yaml:"search_volume"`

	Competition int `json:"competition" yaml:"competition"`

	Opportunity string `json:"opportunity" yaml:"opportunity"`

	Recommendation string `json:"recommendation" yaml:"recommendation"`

	Score float64 `json:"score" yaml:"score"`

	Version int `json:"version" yaml:"version"`

	CreatedAt time.Time `json:"created_at" yaml:"created_at"`
}
