package repository

type Policy struct {
	Object string `yaml:"object"`

	Action string `yaml:"action"`

	Allowed bool `yaml:"allowed"`

	Reason string `yaml:"reason"`
}
