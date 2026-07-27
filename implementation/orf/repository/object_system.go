package repository

// ObjectSystem is the main entry of AOS.
type ObjectSystem struct {
	Repository *Repository

	Registry *Registry

	Graph *ObjectGraph

	Behavior *BehaviorEngine

	BehaviorService *BehaviorService
}

// NewObjectSystem creates an object system.
func NewObjectSystem(repo *Repository) *ObjectSystem {
	engine := NewBehaviorEngine()

	return &ObjectSystem{
		Repository: repo,

		Registry: NewRegistry(repo),

		Graph: NewObjectGraph(),

		Behavior: engine,

		BehaviorService: NewBehaviorService(engine),
	}
}
