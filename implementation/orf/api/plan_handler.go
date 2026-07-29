package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

// CreatePlanRequest represents plan creation request.
type CreatePlanRequest struct {
	Intent string `json:"intent"`

	Name string `json:"name"`

	Steps []string `json:"steps"`
}

// handleObjectPlan handles object plan API.
func (s *Server) handleObjectPlan(
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
			"/plan",
		)

	switch r.Method {

	case http.MethodPost:

		var request CreatePlanRequest

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

		plan :=
			s.System.Plan.Create(
				name,
				request.Intent,
				request.Name,
				request.Steps,
			)

		writeJSON(
			w,
			plan,
		)

	case http.MethodGet:

		plans :=
			s.System.Plan.Get(
				name,
			)

		writeJSON(
			w,
			plans,
		)

	default:

		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)
	}
}
