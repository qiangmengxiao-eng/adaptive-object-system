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

	GraphStore *GraphStore

	GraphService *GraphService

	EventStore *EventStore

	EventService *EventService
}

// NewObjectSystem creates an object system.
func NewObjectSystem(
	repo *Repository,
) *ObjectSystem {

	behavior :=
		NewBehaviorEngine()

	migration :=
		NewMigrationEngine()

	behaviorService :=
		NewBehaviorService(
			behavior,
		)

	migrationService :=
		NewMigrationService(
			migration,
		)

	graphStore :=
		NewGraphStore(
			repo.FS(),
		)

	graphService :=
		NewGraphService(
			graphStore,
		)

	eventStore :=
		NewEventStore(
			repo.FS(),
		)

	eventService :=
		NewEventService(
			eventStore,
		)

	return &ObjectSystem{

		Repository: repo,

		Registry: NewRegistry(
			repo,
		),

		Graph: NewObjectGraph(),

		Behavior: behavior,

		BehaviorService: behaviorService,

		MigrationEngine: migration,

		MigrationService: migrationService,

		GraphStore: graphStore,

		GraphService: graphService,

		EventStore: eventStore,

		EventService: eventService,
	}
}
