package repository

// ObjectSystem is the main entry of AOS.
type ObjectSystem struct {
	Repository *Repository

	Registry *Registry

	Graph *ObjectGraph

	Behavior *BehaviorEngine
}

// NewObjectSystem creates an object system.
func NewObjectSystem(repo *Repository) *ObjectSystem {
	return &ObjectSystem{
		Repository: repo,

		Registry: NewRegistry(repo),

		Graph: NewObjectGraph(),

		Behavior: NewBehaviorEngine(),
	}
}
