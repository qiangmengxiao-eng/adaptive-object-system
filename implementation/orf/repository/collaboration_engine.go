package repository

import "time"

type CollaborationEngine struct {
	Store *CollaborationStore
}

func NewCollaborationEngine(
	store *CollaborationStore,
) *CollaborationEngine {

	return &CollaborationEngine{
		Store: store,
	}
}

func (c *CollaborationEngine) Request(
	from string,
	to string,
	action string,
) (*Collaboration, error) {

	item :=
		Collaboration{

			From: from,

			To: to,

			Action: action,

			Status: "completed",

			Result: "success",

			Version: 1,

			CreatedAt: time.Now(),
		}

	if c.Store != nil {

		_ =
			c.Store.Append(
				item,
			)
	}

	return &item, nil
}
