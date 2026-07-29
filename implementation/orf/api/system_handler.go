package api

import (
	"net/http"
)

// handleSystemStatus handles /system/status.
func (s *Server) handleSystemStatus(
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

	status :=
		s.System.StatusService.Get()

	writeJSON(
		w,
		map[string]int{

			"objects": status.Objects,

			"runtime": status.Runtime,

			"events": status.Events,

			"behaviors": status.Behaviors,

			"audit": status.Audit,
		},
	)
}
