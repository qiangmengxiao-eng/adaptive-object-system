package repository

type Metrics struct {
	Objects int `json:"objects"`

	Executions int `json:"executions"`

	Knowledge int `json:"knowledge"`

	SuccessRate float64 `json:"success_rate"`
}
