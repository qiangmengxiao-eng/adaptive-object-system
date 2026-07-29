package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

type ExecuteBehaviorRequest struct {
	Behavior string `json:"behavior"`
}

// handleObjectBehaviors executes object behavior.
func (s *Server) handleObjectBehaviors(
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
			"/behaviors",
		)

	var request ExecuteBehaviorRequest

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

	if request.Behavior == "" {

		http.Error(
			w,
			"behavior required",
			http.StatusBadRequest,
		)

		return
	}

	err =
		s.System.BehaviorService.Execute(
			request.Behavior,
			name,
		)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	writeJSON(
		w,
		map[string]string{

			"status": "behavior executed",

			"object": name,

			"behavior": request.Behavior,
		},
	)
}
