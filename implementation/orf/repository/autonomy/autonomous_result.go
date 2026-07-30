package autonomy

type AutonomousResult struct {
	Object string `json:"object"`

	Decision string `json:"decision"`

	Plan interface{} `json:"plan"`

	Execution interface{} `json:"execution"`

	Reflection interface{} `json:"reflection"`

	Success bool `json:"success"`
}
