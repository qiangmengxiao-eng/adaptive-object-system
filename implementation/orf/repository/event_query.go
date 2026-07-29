package repository

type EventQuery struct {
	events []ObjectEvent
}

func NewEventQuery(
	events []ObjectEvent,
) *EventQuery {

	return &EventQuery{
		events: events,
	}
}

func (q *EventQuery) ByObject(
	name string,
) []ObjectEvent {

	result :=
		make([]ObjectEvent, 0)

	for _, event := range q.events {

		if event.Object == name {

			result =
				append(
					result,
					event,
				)
		}
	}

	return result
}

func (q *EventQuery) ByType(
	eventType string,
) []ObjectEvent {

	result :=
		make([]ObjectEvent, 0)

	for _, event := range q.events {

		if event.Type == eventType {

			result =
				append(
					result,
					event,
				)
		}
	}

	return result
}

func (q *EventQuery) Latest(
	count int,
) []ObjectEvent {

	if count <= 0 {

		return []ObjectEvent{}
	}

	if count > len(q.events) {

		count = len(q.events)
	}

	return q.events[len(q.events)-count:]
}
