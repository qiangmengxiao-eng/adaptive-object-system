package repository

// ObjectView represents complete object information.
type ObjectView struct {
	Name string

	Definition ObjectDefinition

	Runtime *RuntimeObject

	Events []ObjectEvent

	Behaviors []ObjectBehavior

	Relations []ObjectRelation

	Audit []AuditRecord
}

// ObjectViewService provides unified object view.
type ObjectViewService struct {
	system *ObjectSystem
}

// NewObjectViewService creates object view service.
func NewObjectViewService(
	system *ObjectSystem,
) *ObjectViewService {

	return &ObjectViewService{

		system: system,
	}
}

// Inspect returns complete object view.
func (s *ObjectViewService) Inspect(
	name string,
) (
	*ObjectView,
	error,
) {

	definition, err :=
		s.system.Repository.ReadObjectDefinition(
			name,
		)

	if err != nil {

		return nil, err
	}

	view :=
		&ObjectView{

			Name: name,

			Definition: *definition,

			Events: make(
				[]ObjectEvent,
				0,
			),

			Behaviors: make(
				[]ObjectBehavior,
				0,
			),

			Relations: make(
				[]ObjectRelation,
				0,
			),

			Audit: make(
				[]AuditRecord,
				0,
			),
		}

	// runtime

	if s.system.Runtime != nil {

		if runtime, ok :=
			s.system.Runtime.Get(
				name,
			); ok {

			view.Runtime =
				runtime

			view.Events =
				append(
					view.Events,
					runtime.Events...,
				)
		}
	}

	// graph relations

	if s.system.GraphService != nil {

		relations, err :=
			s.system.GraphService.QueryRelations(
				name,
			)

		if err == nil {

			view.Relations =
				relations
		}
	}

	// behaviors

	if s.system.Behavior != nil {

		view.Behaviors =
			s.system.Behavior.List()
	}

	// audit

	if s.system.AuditStore != nil {

		records, err :=
			s.system.AuditStore.List()

		if err == nil {

			for _, record := range records {

				if record.Object ==
					name {

					view.Audit =
						append(
							view.Audit,
							record,
						)
				}
			}
		}
	}

	return view, nil
}
