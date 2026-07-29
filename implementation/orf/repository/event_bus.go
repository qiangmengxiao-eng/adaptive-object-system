package repository

type EventHandler func(
	ObjectEvent,
) error

type EventBus struct {
	handlers []EventHandler
}

func NewEventBus() *EventBus {

	return &EventBus{

		handlers: make([]EventHandler, 0),
	}
}

func (b *EventBus) Subscribe(
	handler EventHandler,
) {

	b.handlers =
		append(
			b.handlers,
			handler,
		)
}

func (b *EventBus) Publish(
	event ObjectEvent,
) error {

	if err := event.Validate(); err != nil {

		return err
	}

	for _, handler := range b.handlers {

		if err := handler(event); err != nil {

			return err
		}
	}

	return nil
}
