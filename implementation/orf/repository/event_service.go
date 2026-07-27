package repository

type EventService struct {
	store *EventStore
}

func NewEventService(
	store *EventStore,
) *EventService {

	return &EventService{
		store: store,
	}
}

func (s *EventService) Emit(
	event ObjectEvent,
) error {

	if err := event.Validate(); err != nil {
		return err
	}

	events, err := s.store.Load()

	if err != nil {
		return err
	}

	events = append(
		events,
		event,
	)

	return s.store.Save(
		events,
	)
}

func (s *EventService) List() (
	[]ObjectEvent,
	error,
) {

	return s.store.Load()
}
