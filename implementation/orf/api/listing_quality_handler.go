package api

import (
	"net/http"
	"strings"
)

// handleListingQuality evaluates listing quality.
func (s *Server) handleListingQuality(
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

	if s.System.ListingQuality == nil {

		http.Error(
			w,
			"listing quality unavailable",
			http.StatusInternalServerError,
		)

		return
	}

	object :=
		strings.TrimPrefix(
			r.URL.Path,
			"/objects/",
		)

	object =
		strings.TrimSuffix(
			object,
			"/listing/quality",
		)

	listings, err :=
		s.System.ListingStore.Load()

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	for _, item := range listings {

		if item.Object == object {

			result :=
				s.System.ListingQuality.Analyze(
					item,
				)

			writeJSON(
				w,
				result,
			)

			return
		}
	}

	http.Error(
		w,
		"listing not found",
		http.StatusNotFound,
	)
}
