package reflection

type Reflection struct {
	Object string `json:"object"`

	Execution string `json:"execution"`

	Success bool `json:"success"`

	Problem string `json:"problem"`

	Reason string `json:"reason"`

	Learning string `json:"learning"`

	Confidence float64 `json:"confidence"`

	Version int `json:"version"`
}

func NewReflection(
	object string,
	execution string,
) *Reflection {

	return &Reflection{

		Object: object,

		Execution: execution,

		Version: 1,
	}
}
