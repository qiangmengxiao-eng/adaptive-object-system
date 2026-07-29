package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

type GrowthFeedbackRequest struct {
	Impressions int `json:"impressions"`

	Clicks int `json:"clicks"`

	Orders int `json:"orders"`

	Revenue float64 `json:"revenue"`

	AdCost float64 `json:"ad_cost"`
}

// handleGrowthFeedback handles Amazon performance learning.
func (s *Server) handleGrowthFeedback(
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

	if s.System.PerformanceEngine == nil {

		http.Error(
			w,
			"performance engine unavailable",
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
			"/growth-feedback",
		)

	var request GrowthFeedbackRequest

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

	performance :=
		s.System.PerformanceEngine.Analyze(
			object,
			request.Impressions,
			request.Clicks,
			request.Orders,
			request.Revenue,
			request.AdCost,
		)

	reflection :=
		s.System.PerformanceReflection.Reflect(
			performance,
		)

	learning :=
		s.System.GrowthLearning.Learn(
			performance,
		)

	writeJSON(
		w,
		map[string]interface{}{

			"performance": performance,

			"reflection": reflection,

			"learning": learning,
		},
	)
}
