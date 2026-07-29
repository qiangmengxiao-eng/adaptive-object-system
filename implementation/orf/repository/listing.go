package repository

import "time"

// AmazonListing represents generated Amazon product listing.
type AmazonListing struct {
	Object string `yaml:"object" json:"object"`

	Product string `yaml:"product" json:"product"`

	Features []string `yaml:"features" json:"features"`

	Title string `yaml:"title" json:"title"`

	BulletPoints []string `yaml:"bullet_points" json:"bullet_points"`

	Description string `yaml:"description" json:"description"`

	Keywords []string `yaml:"keywords" json:"keywords"`

	Strategy string `yaml:"strategy" json:"strategy"`

	Version int `yaml:"version" json:"version"`

	CreatedAt time.Time `yaml:"created_at" json:"created_at"`
}
