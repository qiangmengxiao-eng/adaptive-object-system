package repository

// Experience summarizes object history.
type Experience struct {
	Object string `yaml:"object"`

	Events int `yaml:"events"`

	Success int `yaml:"success"`

	Failure int `yaml:"failure"`
}

func (e Experience) SuccessRate() float64 {

	if e.Events == 0 {

		return 0
	}

	return float64(e.Success) /
		float64(e.Events)
}
