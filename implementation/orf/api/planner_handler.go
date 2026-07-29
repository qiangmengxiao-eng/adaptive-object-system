package api

import (
	"net/http"
	"strings"
)

// handleObjectGeneratedPlans handles generated plans query.
func (s *Server) handleObjectGeneratedPlans(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodGet {

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
			"/generated-plans",
		)

	if s.System.Planner == nil {

		http.Error(
			w,
			"planner unavailable",
			http.StatusInternalServerError,
		)

		return
	}

	plans :=
		s.System.Planner.Get(
			name,
		)

	writeJSON(
		w,
		plans,
	)
}
