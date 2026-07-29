package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

// CreateIntentRequest represents intent creation request.
type CreateIntentRequest struct {
	Goal string `json:"goal"`

	Name string `json:"name"`

	Purpose string `json:"purpose"`
}

// handleObjectIntent handles object intent API.
func (s *Server) handleObjectIntent(
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
			"/intent",
		)

	switch r.Method {

	case http.MethodPost:

		var request CreateIntentRequest

		err :=
			json.NewDecoder(
				r.Body,
			).Decode(
				&request,
			)

		if err != nil {

			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest,
			)

			return
		}

		intent :=
			s.System.Intent.Create(
				name,
				request.Goal,
				request.Name,
				request.Purpose,
			)

		writeJSON(
			w,
			intent,
		)

	case http.MethodGet:

		intents :=
			s.System.Intent.Get(
				name,
			)

		writeJSON(
			w,
			intents,
		)

	default:

		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)
	}
}
