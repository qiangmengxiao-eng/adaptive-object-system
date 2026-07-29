package api

import (
	"net/http"
	"strings"
)

// handleObjectExperience handles experience query.
func (s *Server) handleObjectExperience(
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
			"/experience",
		)

	experience :=
		s.System.ExperienceEngine.Analyze(
			name,
		)

	writeJSON(
		w,
		map[string]interface{}{

			"object": experience.Object,

			"events": experience.Events,

			"success": experience.Success,

			"failure": experience.Failure,

			"success_rate": experience.SuccessRate(),
		},
	)
}
