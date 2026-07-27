package repository

// ObjectDefinition describes an object definition.
type ObjectDefinition struct {
	Name string `yaml:"name"`

	Type string `yaml:"type,omitempty"`
}
