package api

import (
	"net/http"
	"strings"
)

// handleObjectDecision handles /objects/{name}/decision.
func (s *Server) handleObjectDecision(
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
			"/decision",
		)

	if name == "" {

		http.Error(
			w,
			"object name required",
			http.StatusBadRequest,
		)

		return
	}

	if s.System.Decision == nil {

		http.Error(
			w,
			"decision engine unavailable",
			http.StatusInternalServerError,
		)

		return
	}

	decision, err :=
		s.System.Decision.Decide(
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
		decision,
	)
}
