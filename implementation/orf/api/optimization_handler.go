package api

import (
	"net/http"
	"strings"
)

// handleObjectOptimization handles
// GET /objects/{name}/optimization
func (s *Server) handleObjectOptimization(
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
			"/optimization",
		)

	if name == "" {

		http.Error(
			w,
			"object name required",
			http.StatusBadRequest,
		)

		return
	}

	if s.System.Optimization == nil {

		http.Error(
			w,
			"optimization engine unavailable",
			http.StatusInternalServerError,
		)

		return
	}

	result :=
		s.System.Optimization.Optimize(
			name,
		)

	writeJSON(
		w,
		map[string]interface{}{

			"object": name,

			"optimization": result,
		},
	)
}
