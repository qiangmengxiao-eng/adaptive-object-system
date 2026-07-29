package api

import (
	"net/http"
	"strings"
)

// handleObjectExecutions handles automatic execution records.
func (s *Server) handleObjectExecutions(
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
			"/executions",
		)

	if s.System.AutoExecutor == nil {

		http.Error(
			w,
			"auto executor unavailable",
			http.StatusInternalServerError,
		)

		return
	}

	records :=
		s.System.AutoExecutor.Get(
			name,
		)

	writeJSON(
		w,
		records,
	)
}
