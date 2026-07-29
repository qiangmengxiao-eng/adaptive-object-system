package repository

import "time"

type Collaboration struct {
	From string `yaml:"from"`

	To string `yaml:"to"`

	Action string `yaml:"action"`

	Status string `yaml:"status"`

	Result string `yaml:"result"`

	Version int `yaml:"version"`

	CreatedAt time.Time `yaml:"created_at"`
}
