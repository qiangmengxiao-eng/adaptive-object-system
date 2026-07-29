package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

// CollaborationRequest represents collaboration request.
type CollaborationRequest struct {
	Target string `json:"target"`

	Action string `json:"action"`
}

// handleObjectCollaborate handles
// POST /objects/{name}/collaborate
func (s *Server) handleObjectCollaborate(
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
			"/collaborate",
		)

	var request CollaborationRequest

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

	if s.System.Collaboration == nil {

		http.Error(
			w,
			"collaboration engine unavailable",
			http.StatusInternalServerError,
		)

		return
	}

	result, err :=
		s.System.Collaboration.Request(
			name,
			request.Target,
			request.Action,
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
		result,
	)
}
