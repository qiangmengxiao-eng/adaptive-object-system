package repository

import "fmt"

// ObjectRelation describes a relation between objects.
type ObjectRelation struct {
	From string `yaml:"from"`

	To string `yaml:"to"`

	Type string `yaml:"type"`
}

// Validate checks relation validity.
func (r *ObjectRelation) Validate() error {
	if r == nil {
		return fmt.Errorf("object relation is nil")
	}

	if r.From == "" {
		return fmt.Errorf("relation source is required")
	}

	if r.To == "" {
		return fmt.Errorf("relation target is required")
	}

	if r.Type == "" {
		return fmt.Errorf("relation type is required")
	}

	return nil
}
