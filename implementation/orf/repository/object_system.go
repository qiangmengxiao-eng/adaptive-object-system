package repository

// ObjectSystem is the main entry of AOS.
type ObjectSystem struct {
	Repository *Repository

	Registry *Registry

	Graph *ObjectGraph

	Behavior *BehaviorEngine

	BehaviorService *BehaviorService

	MigrationEngine *MigrationEngine

	MigrationService *MigrationService
}

// NewObjectSystem creates an object system.
func NewObjectSystem(repo *Repository) *ObjectSystem {

	behavior := NewBehaviorEngine()

	migration := NewMigrationEngine()

	// default migration path
	_ = migration.Register(Migration{
		FromVersion: 1,
		ToVersion:   2,
	})

	return &ObjectSystem{
		Repository: repo,

		Registry: NewRegistry(repo),

		Graph: NewObjectGraph(),

		Behavior: behavior,

		BehaviorService: NewBehaviorService(
			behavior,
		),

		MigrationEngine: migration,

		MigrationService: NewMigrationService(
			migration,
		),
	}
}
