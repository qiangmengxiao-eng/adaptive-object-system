package api

import (
	"net/http"
)

// handleSystemMetrics handles /system/metrics.
func (s *Server) handleSystemMetrics(
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

	if s.System.Metrics == nil {

		http.Error(
			w,
			"metrics engine unavailable",
			http.StatusInternalServerError,
		)

		return
	}

	metrics :=
		s.System.Metrics.Collect()

	writeJSON(
		w,
		metrics,
	)
}
