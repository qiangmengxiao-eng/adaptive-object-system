package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

type UpdateStateRequest struct {
	Status string `json:"status"`
}

// handleObjectState updates runtime state.
func (s *Server) handleObjectState(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodPut {

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
			"/state",
		)

	var request UpdateStateRequest

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

	object, ok :=
		s.System.Runtime.Get(
			name,
		)

	if !ok {

		http.Error(
			w,
			"runtime object not found",
			http.StatusNotFound,
		)

		return
	}

	object.State.Status =
		request.Status

	err =
		s.System.Runtime.Save(
			object,
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

			"status": "state updated",

			"object": name,

			"state": request.Status,
		},
	)
}
