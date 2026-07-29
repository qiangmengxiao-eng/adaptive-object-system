package api

import (
	"net/http"
	"strings"
)

// handleListingOptimize optimizes listing.
func (s *Server) handleListingOptimize(
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

	object :=
		strings.TrimPrefix(
			r.URL.Path,
			"/objects/",
		)

	object =
		strings.TrimSuffix(
			object,
			"/listing/optimize",
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

			result, quality, err :=
				s.System.ListingOptimizer.Optimize(
					item,
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
				map[string]interface{}{

					"listing": result,

					"quality": quality,
				},
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
