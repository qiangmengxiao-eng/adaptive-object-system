package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

type CreateGoalRequest struct {
	Name string `json:"name"`

	Description string `json:"description"`
}

// handleObjectGoal handles object goal API.
func (s *Server) handleObjectGoal(
	w http.ResponseWriter,
	r *http.Request,
) {

	name :=
		strings.TrimPrefix(
			r.URL.Path,
			"/objects/",
		)

	name =
		strings.TrimSuffix(
			name,
			"/goal",
		)

	switch r.Method {

	case http.MethodPost:

		var request CreateGoalRequest

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

		goal :=
			s.System.Goal.Create(
				name,
				request.Name,
				request.Description,
			)

		writeJSON(
			w,
			goal,
		)

	case http.MethodGet:

		goals :=
			s.System.Goal.Get(
				name,
			)

		writeJSON(
			w,
			goals,
		)

	default:

		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)
	}
}
