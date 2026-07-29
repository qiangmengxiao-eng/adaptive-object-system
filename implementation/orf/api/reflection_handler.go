package api

import (
	"net/http"
	"strings"
)

// handleObjectReflection handles reflection query.
func (s *Server) handleObjectReflection(
	w http.ResponseWriter,
	r *http.Request,
) {

	name :=
		strings.TrimSuffix(
			strings.TrimPrefix(
				r.URL.Path,
				"/objects/",
			),
			"/reflection",
		)

	switch r.Method {

	case http.MethodGet:

		if s.System.Reflection == nil {

			writeJSON(
				w,
				nil,
			)

			return
		}

		reflection :=
			s.System.Reflection.Latest(
				name,
			)

		writeJSON(
			w,
			reflection,
		)

	default:

		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)
	}
}
