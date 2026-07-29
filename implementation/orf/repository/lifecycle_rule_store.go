package repository

import (
	"gopkg.in/yaml.v3"
)

const LifecycleRuleFileName = "rules.yaml"

// LifecycleRuleStore persists lifecycle rules.
type LifecycleRuleStore struct {
	fs Storage
}

// NewLifecycleRuleStore creates lifecycle rule store.
func NewLifecycleRuleStore(
	fs Storage,
) *LifecycleRuleStore {

	return &LifecycleRuleStore{
		fs: fs,
	}
}

// Save persists lifecycle rules.
func (s *LifecycleRuleStore) Save(
	rules []LifecycleRule,
) error {

	data, err :=
		yaml.Marshal(
			rules,
		)

	if err != nil {
		return err
	}

	return s.fs.WriteFile(
		"/"+LifecycleRuleFileName,
		data,
	)
}

// Load restores lifecycle rules.
func (s *LifecycleRuleStore) Load() (
	[]LifecycleRule,
	error,
) {

	data, err :=
		s.fs.ReadFile(
			"/" + LifecycleRuleFileName,
		)

	if err != nil {

		return []LifecycleRule{}, nil
	}

	var rules []LifecycleRule

	err =
		yaml.Unmarshal(
			data,
			&rules,
		)

	if err != nil {

		return nil, err
	}

	return rules, nil
}
