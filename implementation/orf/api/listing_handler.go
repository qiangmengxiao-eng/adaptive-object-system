package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

type ListingRequest struct {
	Product string `json:"product"`

	Features []string `json:"features"`

	Keywords []string `json:"keywords"`
}

// handleObjectListing generates Amazon listing.
func (s *Server) handleObjectListing(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodPost {

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
			"/listing",
		)

	var request ListingRequest

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

	if s.System.Listing == nil {

		http.Error(
			w,
			"listing engine unavailable",
			http.StatusInternalServerError,
		)

		return
	}

	result, err :=
		s.System.Listing.Generate(
			name,
			request.Product,
			request.Features,
			request.Keywords,
		)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	writeJSON(
		w,
		result,
	)
}
