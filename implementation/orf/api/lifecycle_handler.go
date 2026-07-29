package api

import (
	"net/http"
	"strings"
)

// handleObjectLifecycle handles
// GET /objects/{name}/lifecycle
func (s *Server) handleObjectLifecycle(
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
			"/lifecycle",
		)

	if name == "" {

		http.Error(
			w,
			"object name required",
			http.StatusBadRequest,
		)

		return
	}

	if s.System.LifecycleManager == nil {

		http.Error(
			w,
			"lifecycle manager unavailable",
			http.StatusInternalServerError,
		)

		return
	}

	state :=
		s.System.LifecycleManager.Evaluate(
			name,
		)

	writeJSON(
		w,
		map[string]interface{}{

			"object": name,

			"state": state,
		},
	)
}
