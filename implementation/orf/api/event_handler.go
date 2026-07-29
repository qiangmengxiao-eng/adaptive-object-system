package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/repository"
)

type CreateEventRequest struct {
	Type string `json:"type"`

	Action string `json:"action"`
}

func (s *Server) handleObjectEvents(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodPost {

		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)

		return
	}

	name :=
		strings.TrimPrefix(
			r.URL.Path,
			"/objects/",
		)

	name =
		strings.TrimSuffix(
			name,
			"/events",
		)

	var request CreateEventRequest

	err :=
		json.NewDecoder(
			r.Body,
		).Decode(
			&request,
		)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	event :=
		repository.NewObjectEvent(
			request.Type,
			name,
			request.Action,
			"",
		)

	err =
		s.System.EventBus.Publish(
			event,
		)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	_ =
		s.System.Runtime.AddEvent(
			name,
			event,
		)

	// Adaptive Observation

	result := "success"

	if strings.Contains(
		request.Type,
		"failed",
	) {

		result = "failure"
	}

	observation :=
		repository.NewObservation(
			name,
			request.Type,
			request.Action,
			result,
		)

	_ =
		s.System.ObservationStore.Append(
			observation,
		)

	// Decision -> Adaptation pipeline
	//
	// Must execute after observation is stored.
	// Otherwise decision reads old experience.

	if s.System.Decision != nil &&
		s.System.Adaptation != nil {

		decision, err :=
			s.System.Decision.Decide(
				name,
			)

		if err == nil &&
			decision != nil {

			_,
				_ =
				s.System.Adaptation.Adapt(
					name,
					decision,
				)
		}
	}

	// Lifecycle transition

	object, ok :=
		s.System.Runtime.Get(
			name,
		)

	if ok {

		_ =
			object
	}

	writeJSON(
		w,
		map[string]string{

			"status": "event published",

			"object": name,
		},
	)
}
