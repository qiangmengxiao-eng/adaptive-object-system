package repository

// ObjectSystem is the main entry of AOS.
type ObjectSystem struct {
	Repository *Repository

	Registry *Registry

	ObjectViewService *ObjectViewService

	Lifecycle *LifecycleEngine

	Graph *ObjectGraph

	GraphStore *GraphStore

	GraphService *GraphService

	Behavior *BehaviorEngine

	BehaviorService *BehaviorService

	MigrationEngine *MigrationEngine

	MigrationService *MigrationService

	EventStore *EventStore

	EventService *EventService

	EventBus *EventBus

	Runtime *RuntimeEngine

	RuntimeStore *RuntimeStore

	RuleEngine *RuleEngine

	RuleService *RuleService

	AuditStore *AuditStore

	StatusService *StatusService

	QueryService *QueryService

	// Adaptive Intelligence

	ObservationStore *ObservationStore

	ExperienceEngine *ExperienceEngine

	Decision *DecisionEngine

	Adaptation *AdaptationEngine

	Goal *GoalEngine

	Intent *IntentEngine

	Plan *PlanEngine

	Task *TaskEngine

	Planner *PlannerEngine

	AutoExecutor *AutoExecutor

	ExecutionStore *ExecutionStore

	ReflectionStore *ReflectionStore

	Reflection *ReflectionEngine

	KnowledgeStore *KnowledgeStore

	Knowledge *KnowledgeEngine

	ListingStore *ListingStore

	Listing *ListingEngine

	ListingQuality *ListingQualityEngine

	ListingQualityStore *ListingQualityStore

	ListingOptimizer *ListingOptimizer

	// Phase 4

	Collaboration *CollaborationEngine

	CollaborationStore *CollaborationStore

	LifecycleManager *LifecycleManager

	Policy *PolicyEngine

	Metrics *MetricsEngine

	Optimization *OptimizationEngine
}

// NewObjectSystem creates object system.
func NewObjectSystem(
	repo *Repository,
) *ObjectSystem {

	behavior :=
		NewBehaviorEngine()

	behaviorService :=
		NewBehaviorService(
			behavior,
		)

	_ =
		behavior.Register(
			ObjectBehavior{

				Name: "initialize",

				Action: "activate",
			},
		)

	migration :=
		NewMigrationEngine()

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

	eventBus :=
		NewEventBus()

	runtimeStore :=
		NewRuntimeStore(
			repo.FS(),
		)

	runtime :=
		NewRuntimeEngine(
			runtimeStore,
		)

	ruleEngine :=
		NewRuleEngine()

	ruleService :=
		NewRuleService(
			ruleEngine,
			behavior,
		)

	lifecycleStore :=
		NewLifecycleRuleStore(
			repo.FS(),
		)

	lifecycle :=
		NewLifecycleEngine(
			lifecycleStore,
		)

	auditStore :=
		NewAuditStore(
			repo.FS(),
		)

	// Adaptive Memory

	observationStore :=
		NewObservationStore(
			repo.FS(),
		)

	reflectionStore :=
		NewReflectionStore(
			repo.FS(),
		)

	experienceEngine :=
		NewExperienceEngine(
			observationStore,
		)
	reflection :=
		NewReflectionEngine(
			experienceEngine,
			reflectionStore,
		)

	knowledgeStore :=
		NewKnowledgeStore(
			repo.FS(),
		)
	listingStore :=
		NewListingStore(
			repo.FS(),
		)

	knowledge :=
		NewKnowledgeEngine(
			experienceEngine,
			knowledgeStore,
		)

	listing :=
		NewListingEngine(
			knowledge,
			listingStore,
		)

	listingQualityStore :=
		NewListingQualityStore(
			repo.FS(),
		)

	listingQuality :=
		NewListingQualityEngine()

	listingOptimizer :=
		NewListingOptimizer(
			listing,
			listingQuality,
			knowledge,
		)

	// Decision Engine

	decision :=
		NewDecisionEngine(
			experienceEngine,
		)

	decision.AttachReflection(
		reflection,
	)

	decision.AttachKnowledge(
		knowledge,
	)

	adaptation :=
		NewAdaptationEngine(
			experienceEngine,
			behavior,
		)

	adaptation.AttachKnowledge(
		knowledge,
	)

	goal :=
		NewGoalEngine()

	intent :=
		NewIntentEngine()

	plan :=
		NewPlanEngine()

	task :=
		NewTaskEngine()

	planner :=
		NewPlannerEngine()

	executionStore :=
		NewExecutionStore(
			repo.FS(),
		)

	autoExecutor :=
		NewAutoExecutor(
			task,
			executionStore,
		)

	autoExecutor.Load()

	// Phase 4 Autonomous Layer

	collaborationStore :=
		NewCollaborationStore(
			repo.FS(),
		)

	collaboration :=
		NewCollaborationEngine(
			collaborationStore,
		)

	lifecycleManager :=
		NewLifecycleManager(
			experienceEngine,
		)

	policy :=
		NewPolicyEngine()

	optimization :=
		NewOptimizationEngine(
			decision,
		)
	_ =
		ruleEngine.Register(
			ObjectRule{

				Name: "initialize-object",

				Event: "object.created",

				Action: "initialize",
			},
		)

	system :=
		&ObjectSystem{

			Repository: repo,

			Adaptation: adaptation,

			Goal: goal,

			Intent: intent,

			Plan: plan,

			Task: task,

			Planner: planner,

			AutoExecutor: autoExecutor,

			ExecutionStore: executionStore,

			ObservationStore: observationStore,

			ExperienceEngine: experienceEngine,

			Decision: decision,

			ReflectionStore: reflectionStore,

			Reflection: reflection,

			KnowledgeStore: knowledgeStore,

			Knowledge: knowledge,

			ListingStore: listingStore,

			Listing: listing,

			ListingQualityStore: listingQualityStore,

			ListingQuality: listingQuality,

			ListingOptimizer: listingOptimizer,

			// Phase 4

			CollaborationStore: collaborationStore,

			Collaboration: collaboration,

			LifecycleManager: lifecycleManager,

			Policy: policy,

			Optimization: optimization,

			Lifecycle: lifecycle,

			Graph: NewObjectGraph(),

			GraphStore: graphStore,

			GraphService: graphService,

			Behavior: behavior,

			BehaviorService: behaviorService,

			MigrationEngine: migration,

			MigrationService: migrationService,

			EventStore: eventStore,

			EventService: eventService,

			EventBus: eventBus,

			Runtime: runtime,

			RuntimeStore: runtimeStore,

			RuleEngine: ruleEngine,

			RuleService: ruleService,

			AuditStore: auditStore,
		}

	system.Registry =
		NewRegistry(
			repo,
		)

	system.Metrics =
		NewMetricsEngine(
			system.Registry,
			system.ExecutionStore,
			system.KnowledgeStore,
			system.ObservationStore,
		)

	system.ObjectViewService =
		NewObjectViewService(
			system,
		)

	system.StatusService =
		NewStatusService(
			system,
		)

	system.QueryService =
		NewQueryService(
			system.Registry,
		)

	// Event pipeline

	eventBus.Subscribe(
		func(
			event ObjectEvent,
		) error {

			_ =
				eventService.Emit(
					event,
				)

			rules :=
				ruleService.Handle(
					event,
				)

			for _, rule := range rules {

				if rule.Action !=
					"initialize" {

					continue
				}

				_ =
					behaviorService.Execute(
						rule.Action,
						event.Object,
					)

				if object, ok :=
					runtime.Get(
						event.Object,
					); ok {

					object.State.Status =
						"active"

					executedEvent :=
						NewObjectEvent(
							"behavior.executed",
							event.Object,
							"initialize",
							"",
						)

					_ =
						runtime.AddEvent(
							event.Object,
							executedEvent,
						)

					_ =
						eventService.Emit(
							executedEvent,
						)

					_ =
						auditStore.Append(
							NewAuditRecord(
								"behavior.initialize",
								event.Object,
								"success",
							),
						)
				}
			}

			return nil
		},
	)

	return system
}
