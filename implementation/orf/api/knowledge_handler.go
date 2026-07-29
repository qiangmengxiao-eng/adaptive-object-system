package api

import (
	"net/http"
	"strings"
)

func (s *Server) handleObjectKnowledge(
	w http.ResponseWriter,
	r *http.Request,
) {

	name :=
		strings.TrimPrefix(
			r.URL.Path,
			"/objects/",
		)

	name =
		strings.TrimSuffix(
			name,
			"/knowledge",
		)

	result :=
		s.System.Knowledge.Get(
			name,
		)

	writeJSON(
		w,
		result,
	)
}
