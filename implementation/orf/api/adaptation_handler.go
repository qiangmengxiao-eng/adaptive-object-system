package api

import (
	"net/http"
	"strings"
)

// handleObjectAdaptations handles /objects/{name}/adaptations.
func (s *Server) handleObjectAdaptations(
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
			"/adaptations",
		)

	if name == "" {

		http.Error(
			w,
			"object name required",
			http.StatusBadRequest,
		)

		return
	}

	if s.System.Adaptation == nil {

		http.Error(
			w,
			"adaptation engine unavailable",
			http.StatusInternalServerError,
		)

		return
	}

	writeJSON(
		w,
		s.System.Adaptation.List(
			name,
		),
	)
}
