package repository

// SystemStatus describes system metrics.
type SystemStatus struct {
	Objects int

	Runtime int

	Events int

	Behaviors int

	Audit int
}

// StatusService provides system status.
type StatusService struct {
	system *ObjectSystem
}

// NewStatusService creates status service.
func NewStatusService(
	system *ObjectSystem,
) *StatusService {

	return &StatusService{

		system: system,
	}
}

// Get returns system status.
func (s *StatusService) Get() SystemStatus {

	status :=
		SystemStatus{}

	objects, err :=
		s.system.Registry.List()

	if err == nil {

		status.Objects =
			len(objects)
	}

	if s.system.Runtime != nil {

		status.Runtime =
			len(
				s.system.Runtime.List(),
			)
	}

	if s.system.EventService != nil {

		events, err :=
			s.system.EventService.List()

		if err == nil {

			status.Events =
				len(events)
		}
	}

	if s.system.Behavior != nil {

		status.Behaviors =
			len(
				s.system.Behavior.List(),
			)
	}

	// audit count

	if s.system.AuditStore != nil {

		records, err :=
			s.system.AuditStore.List()

		if err == nil {

			status.Audit =
				len(records)
		}
	}

	return status
}
